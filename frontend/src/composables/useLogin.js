import {
    ref
} from "vue"

const error = ref(null)

const login = async (creds, password) => {
    error.value = null
    let data = {
        creds: creds,
        password: password
    }
    console.log(JSON.stringify(data))
    try {
        let res = await fetch("./api/login", {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(data)
        })
        if (!res.ok) {
            throw Error("Failed to login")
        }
    } catch (err) {
        error.value = err.message
        console.log(err.message)
    }
}

const useLogin = () => {
    return {
        error,
        login
    }
}

export default useLogin