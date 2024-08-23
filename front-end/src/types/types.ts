export type Question = {
  id: number;
  content: string;
  created_at: string;
  answered: boolean;
  answer: string[] | null;
};

export type User = {
  username: string;
  email: string;
  questions: Question[];
  created_at: string;
};
