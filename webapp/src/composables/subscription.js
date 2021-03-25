import {
    ref
} from "@vue/reactivity"



const useSubscription = () => {
    const error = ref(null)
    const followersList = ref([])
    const isFollowing = ref(false)
    const follow = async (followingId) => {
        let data = {
            following_id: followingId,
        }
        try {
            let res = await fetch('./api/follow', {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(data)
            })
    
            if (!res.ok) {
                throw Error("Failed to follow")
            }
            isFollowing.value = true
        } catch (e) {
            console.log(e.message)
            error.value = e.message
        }
    }
    
    const unfollow = async (followingId) => {
        let data = {
            following_id: followingId,
        }
        try {
            let res = await fetch('./api/unfollow', {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(data)
            })
    
            if (!res.ok) {
                throw Error("Failed to unfollow")
            }
            isFollowing.value = false
        } catch (e) {
            console.log(e.message)
            error.value = e.message
        }
    }
    
    const checkIfFollowing = async (followingId) => {
        try {
            let res = await fetch(`./api/isfollowing?following_id=${followingId}`)
            if (!res.ok) {
                throw Error("Failed to check if following")
            }
            let data = await res.json()
            console.log(data)
            isFollowing.value = data.is_following
        } catch (e) {
            console.log(e.message)
            error.value = e.message
        }
    }
    const getFollowers = async (profileId) => {
        try {
            let res = await fetch(`./api/followers?profile_id=${profileId}`)
            if (!res.ok) {
                throw Error("Failed to check of following")
            }
            let data = await res.json()
            followersList.value = data.followers_list
            console.log(followersList.value)
        } catch (e) {
            console.log(e.message)
            error.value = e.message
        }
    }
    return {
        error,
        followersList,
        isFollowing,
        checkIfFollowing,
        follow,
        unfollow,
        getFollowers,
    }
}

export default useSubscription