import cookie from 'react-cookies'

// get cookie of current user
export const loginUser = () => {
    return cookie.load('userInfo')
}

// user login, save cookie
// export const onLogin = (user) =>{
//     cookie.save('userInfo', user, { path: '/' })
// }
export const onLogin = (info) =>{
    cookie.save('userInfo', {'id': info}, { path: '/' })
}

// user logout, remove cookie
export const logout = () => {
    cookie.remove('userInfo')
    window.location.href = '/login'
}
