import { ClientConfiguration } from 'twirpscript/dist'
import { Login, Logout, ValidSites } from './pb/access.pb'
import personApi from './person'

const handler: ProxyHandler<any> = {
  apply: function (target: any, thisArg: any, argumentsList: any[]) {
    let payload = null
    if (argumentsList.length > 0) {
      payload = argumentsList[0]
    }
    let config: ClientConfiguration = {}
    if (argumentsList.length > 1) {
      config = {
        ...config,
        ...argumentsList[1],
      }
    }
    return target(payload, config)
  },
}

const api = {
  access: {
    Login: new Proxy(Login, handler),
    Logout: Logout,
    ValidSites: ValidSites,
  },
  person: personApi
}

export default api
