export interface Post {
  title: string;
  summary: string;
  body: string;
  authorId: number;
  published?: boolean;
  dateCreated?: string;
  dateUpdated?: string;
  id?: number;
  color?: string;
  img?: string;
}

export interface PostsState {
  posts?: Post[];
  error: boolean;
}
