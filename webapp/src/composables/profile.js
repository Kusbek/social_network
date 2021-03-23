import {
    ref
} from "@vue/reactivity"
const profile = ref(null)
const error = ref(null)
const load = async (id) => {
    try {
        let res = await fetch(`./api/user?id=${id}`)
        let data = await res.json()
        profile.value = {
            ...data
        }
    } catch (e) {
        console.log(e.message)
        error.value = "Could not fetch profile"
    }
    // profile.value = {
    //     id: id,
    //     username: "kusbek_test",
    //     email: "kusbek1994_test@gmail.com",
    //     firstName: "Bekarys_test",
    //     lastName: "Kuspan_test",
    //     birthDate: "1994-09-18",
    //     aboutMe: "Test testinson motherfucking testicles for mistestress",
    //     pathToPhoto: "/img/images/avatars/default.jpg"
    // }
}

const follow = async (followingId) => {
    let data = {
        following_id: followingId,
    }
    try {
        let res = await fetch('./api/user/follow', {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(data)
        })
        
        if (!res.ok) {
            throw Error("Failed to follow")
        }
        profile.value.isFollowing = true
    } catch (e) {
        console.log(e.message)
        error.value = e.message
    }
}

const unfollow = async () => {
    console.log("unfollow")
    // profile.value.isFollowing = false
    console.log(profile.value)
}

const useProfile = () => {
    return {
        profile,
        error,
        follow,
        unfollow,
        load,
    }
}

export default useProfile