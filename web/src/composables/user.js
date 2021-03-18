import {
    ref
} from "@vue/reactivity"


const error = ref(null)
const user = ref(null)

const getUser = async () => {
    if (!user.value) {
        await fetch("./api/auth").then(res => {
            if (!res.ok || res.status == 401) {
                throw Error("Failed to authenticate")
            }
            return res.json()
        }).then(data => {
            user.value = {
                ...data
            }
        }).catch (e => {
            user.value = null
            console.log(e.message)
        }) 
    }
}

const signup = async (username, email, firstName, lastName, birthDate, aboutMe, password) => {
    error.value = null
    user.value = null
    let body = {
        username: username,
        email: email,
        first_name: firstName,
        last_name: lastName,
        birth_date: birthDate,
        about_me: aboutMe,
        password: password,
    }
    try {
        let res = await fetch("./api/signup", {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(body)
        })
        if (!res.ok) {
            throw Error("Failed to sign up")
        }
        await getUser()
    } catch (err) {
        error.value = err.message
        console.log(err.message)
    }
}

const logout = async () => {
    error.value = null
    try {
        let res = await fetch("./api/logout", {
            method: "POST",
            headers: {"Content-Type": "application/json"},
        })
        if (!res.ok) {
            throw Error("Failed to logout")
        }
        user.value = null
    } catch (err) {
        error.value = err.message
    }
}

const login = async (creds, password) => {
    error.value = null
    user.value = null
    let data = {
        creds: creds,
        password: password
    }
    try {
        let res = await fetch("./api/login", {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(data)
        })
        if (!res.ok) {
            throw Error("Failed to login")
        }
        await getUser()
    } catch (err) {
        error.value = err.message
        console.log(err.message)
    }
}

const User = () => {
    return {
        error,
        user,
        getUser,
        signup,
        login,
        logout,
    }
}

export default User