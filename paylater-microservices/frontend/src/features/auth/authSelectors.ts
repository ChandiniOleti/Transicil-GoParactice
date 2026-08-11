import type { RootState } from '../../app/store'

export const selectAuthState = (state: RootState) => state.auth

export const selectToken = (state: RootState) => state.auth.token

export const selectCurrentUser = (state: RootState) => state.auth.user

export const selectAuthRole = (state: RootState) => state.auth.role

export const selectAuthStatus = (state: RootState) => state.auth.status

export const selectIsAuthenticated = (state: RootState) =>
  state.auth.isAuthenticated

export const selectAuthLoading = (state: RootState) => state.auth.isLoading

export const selectAuthError = (state: RootState) => state.auth.error
