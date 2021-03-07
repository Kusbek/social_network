import {
    ref
} from "vue";

const user = ref(null)

const getUser = async () => {
    try {
        let res = await fetch("./api/auth")
        if (!res.ok) {
            throw Error("Failed to authenticate")
        }

        data = await res.json()
        console.log(data)
        user.value =  {
            ...data
        }

    } catch (err) {
        console.log(err.message)
    }
    return {
        user
    }
}

export default getUser