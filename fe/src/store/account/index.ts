import type { Module } from "vuex";
import mutations from "./mutation";
import { AccountState } from "./state";
import type { RootState } from "../rootState";

export const state = new AccountState();

const namespaced = true;

export const account: Module<AccountState, RootState> = {
  namespaced,
  state,
  mutations,
};