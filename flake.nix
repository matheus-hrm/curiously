{
  description = "Ambiente de desenvolvimento em Go com PostgreSQL";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShells.default = pkgs.mkShell
          {
            buildInputs = with pkgs; [
              go
              gotools
              gopls
              go-outline
              gopkgs
              godef
              golint
              air
              postgresql
              templ
              tailwindcss
            ];

            shellHook = ''
              export GOPATH=$HOME/go
              export PATH=$PATH:$GOPATH/bin
              
              # Configuração do PostgreSQL
              export PGDATA=$PWD/postgres_data
              export PGHOST=localhost
              export PGPORT=5433  # Mudamos para 5433 para evitar conflitos

              echo "Criando diretório de dados do PostgreSQL..."
              mkdir -p $PGDATA
              chmod 700 $PGDATA

              if [ ! -f $PGDATA/postgresql.conf ]; then
                echo "Inicializando banco de dados PostgreSQL..."
                initdb -D $PGDATA
                echo "listen_addresses = '*'" >> $PGDATA/postgresql.conf
                echo "port = 5433" >> $PGDATA/postgresql.conf
                echo "unix_socket_directories = '$PWD/postgres_run'" >> $PGDATA/postgresql.conf
                echo "host all all 127.0.0.1/32 trust" >> $PGDATA/pg_hba.conf
              fi

              echo "Iniciando servidor PostgreSQL..."
              mkdir -p postgres_run
              pg_ctl -D $PGDATA -o "-k '$PWD/postgres_run'" -l $PWD/postgres.log start

              # Espera o servidor iniciar
              for i in {1..30}; do
                if pg_isready -h localhost -p 5433; then
                  echo "Servidor PostgreSQL está pronto!"
                  break
                fi
                echo "Aguardando servidor PostgreSQL iniciar... ($i/30)"
                sleep 1
              done

              if ! pg_isready -h localhost -p 5433; then
                echo "Falha ao iniciar o servidor PostgreSQL. Verifique o arquivo postgres.log para mais detalhes."
                cat $PWD/postgres.log
              else
                echo "Ambiente de desenvolvimento Go com PostgreSQL está pronto!"
                echo "Para conectar ao banco de dados: psql -h localhost -p 5433"
              fi
            '';

            # Limpeza ao sair do shell
            exitHook = ''
              echo "Parando servidor PostgreSQL..."
              pg_ctl -D $PGDATA stop
            '';
          };
      }
    );
}
