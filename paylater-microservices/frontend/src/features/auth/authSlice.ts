import { createSlice, type PayloadAction } from '@reduxjs/toolkit'



import type { AuthRole } from '../../types/auth'

import { isTokenExpired, readJwtClaims } from '../../utils/jwt'

import { getToken, removeToken, saveToken } from '../../utils/token'



/**

 * Authenticated identity derived from JWT claims (user_id, email, role).

 */

export interface AuthUser {

  user_id: number

  email: string

  role: AuthRole

}



export type AuthStatus =

  | 'idle'

  | 'loading'

  | 'authenticated'

  | 'unauthenticated'



export interface AuthState {

  token: string | null

  user: AuthUser | null

  role: AuthRole | null

  status: AuthStatus

  isAuthenticated: boolean

  isLoading: boolean

  error: string | null

}



export interface SetCredentialsPayload {

  token: string

  user: AuthUser

}



function createInitialAuthState(): AuthState {

  const unauthenticated: AuthState = {

    token: null,

    user: null,

    role: null,

    status: 'unauthenticated',

    isAuthenticated: false,

    isLoading: false,

    error: null,

  }



  const token = getToken()

  if (!token) {

    return unauthenticated

  }



  const claims = readJwtClaims(token)

  if (!claims || isTokenExpired(claims.exp)) {

    removeToken()

    return unauthenticated

  }



  const user: AuthUser = {

    user_id: claims.user_id,

    email: claims.email,

    role: claims.role,

  }



  return {

    token,

    user,

    role: claims.role,

    status: 'authenticated',

    isAuthenticated: true,

    isLoading: false,

    error: null,

  }

}



const authSlice = createSlice({

  name: 'auth',

  initialState: createInitialAuthState(),

  reducers: {

    setCredentials(state, action: PayloadAction<SetCredentialsPayload>) {

      const { token, user } = action.payload

      state.token = token

      state.user = user

      state.role = user.role

      state.status = 'authenticated'

      state.isAuthenticated = true

      state.isLoading = false

      state.error = null

      saveToken(token)

    },

    logout(state) {

      state.token = null

      state.user = null

      state.role = null

      state.status = 'unauthenticated'

      state.isAuthenticated = false

      state.isLoading = false

      state.error = null

      removeToken()

    },

  },

})



export const { setCredentials, logout } = authSlice.actions

export default authSlice.reducer

