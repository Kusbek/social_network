import {
    ref
} from "@vue/reactivity"



const useSubscription = () => {
    const error = ref(null)
    const followersList = ref([])
    const followingList = ref([])
    const isFollowing = ref(false)
    const follow = async (followingId) => {
        let data = {
            following_id: parseInt(followingId),
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
            following_id: parseInt(followingId),
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
        } catch (e) {
            console.log(e.message)
            error.value = e.message
        }
    }
    const getFollowing = async (profileId) => {
        try {
            let res = await fetch(`./api/following?profile_id=${profileId}`)
            if (!res.ok) {
                throw Error("Failed to check of following")
            }
            let data = await res.json()
            followingList.value = data.following_list
        } catch (e) {
            console.log(e.message)
            error.value = e.message
        }
    }
    return {
        error,
        followersList,
        followingList,
        isFollowing,
        checkIfFollowing,
        follow,
        unfollow,
        getFollowers,
        getFollowing,
    }
}

export default useSubscription