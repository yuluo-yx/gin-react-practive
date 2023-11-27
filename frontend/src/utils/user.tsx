export function getCurUserFromBrowser() {

    return localStorage.getItem("user")
}

export function getCurUserToken() {

    return localStorage.getItem("token")
}
