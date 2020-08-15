import { isValidToken } from '@/utils'

export default function ({ store, redirect }) {
  // if the user is not authenticated
  if (!isValidToken(store.state.user.token)) {
    return redirect('/')
  }
}
