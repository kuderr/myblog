export class User {
  id: number;
  username: string;
  role: string;
  email: string;
  fullName: string;
  blocked: boolean;

  constructor() {
    this.id = 0;
    this.username = "";
    this.role = "";
    this.email = "";
    this.fullName = "";
    this.blocked = false;
  }
}
