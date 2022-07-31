import { Login, Logout, SessionInfo, UnauthorisedUser } from '@/api/pb/access.pb'
import router from '@/router'
import { defineStore } from 'pinia'

interface IUserState {
  isCheckingLogin: boolean
  isValid: boolean
  details: SessionInfo
}

function anon() {
  return {
    token: '',
    expiration: 0,
    user: {
      username: 'Anonymous',
      name: 'Anonymous',
      email: '',
      siteUlid: '',
      ulid: '',
      password: '',
      role: '',
      initials: '',
      picture: '',
    },
    sites: [],
  }
}

export const useUserAccessStore = defineStore('access', {
  state: (): IUserState => ({
    isCheckingLogin: false,
    isValid: false,
    details: anon(),
  }),
  getters: {
    isLoggedIn(state: IUserState): boolean {
      return this.details.user.siteUlid !== '' && this.details.user.ulid !== ''
    },
    siteUlid(state: IUserState): string {
      return state.details.user.siteUlid
    },
  },
  actions: {
    async login(unauthorisedUser: UnauthorisedUser) {
      let sessionInfo = await Login(unauthorisedUser)
      this.details = sessionInfo
      if (this.isLoggedIn) {
        router.push({ name: 'home' })
      }
    },
    logout() {
      // let sessionInfo = Logout(unauthorisedUser)
      Logout({
        email: this.details.user.email,
        siteUlid: this.details.user.siteUlid
      })
      this.details = anon()
    },
  },
  persist: true,
})
