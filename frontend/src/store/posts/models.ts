export class Post {
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

  constructor() {
    this.title = "";
    this.summary = "";
    this.body = "";
    this.authorId = 1;
  }
}
