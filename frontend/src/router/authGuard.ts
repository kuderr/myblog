import store from "@/store";
import { isValidToken } from "@/utils";

export function authRequired(to, from, next) {
  if (!isValidToken(store.state.user.token)) {
    next("/");
  } else {
    next();
  }
}
