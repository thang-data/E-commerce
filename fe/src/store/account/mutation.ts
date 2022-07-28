import type { MutationTree } from "vuex"
import type { AccountState } from "./state"
const mutations: MutationTree<AccountState> = {
  updateFullName(state, account: AccountState) {
    state.id = account.id
    state.firstName = account.firstName
    state.lastName = account.lastName
  },
  updateFullAccount(state, account: AccountState) {
    state.id = account.id
    state.firstName = account.firstName
    state.lastName = account.lastName
  },
}

export default mutations
