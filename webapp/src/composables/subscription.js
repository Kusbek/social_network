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
            let res = await fetch('./api/user/unfollow', {
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
            let res = await fetch(`./api/user/isfollowing?following_id=${followingId}`)
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
            // let res = await fetch(`./api/user/followers`)
            // if (!res.ok) {
            //     throw Error("Failed to check of following")
            // }
            // let data = await res.json()
            let data = {
                followers_list: [{
                        id: 1,
                        first_name: "Bekarys",
                        last_name: "Kuspan",
                        path_to_photo: "/img/ninja.jpg",
                    },
                    {
                        id: 2,
                        first_name: "Scarlett",
                        last_name: "Johanson",
                        path_to_photo: "/img/ninja.jpg",
                    },
                    {
                        id: 3,
                        first_name: "Test",
                        last_name: "User",
                        path_to_photo: "/img/ninja.jpg",
                    },
                ]
            }
            followersList.value = data.followers_list
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