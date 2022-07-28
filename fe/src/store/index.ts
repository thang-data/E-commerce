import Vuex from "vuex";
import type { StoreOptions } from "vuex";
import type { RootState } from "./rootState";
import { account } from "./account";
const store: StoreOptions<RootState> = {
  modules: {
    account
  },
};

export default new Vuex.Store<RootState>(store);

