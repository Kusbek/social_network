import {
    ref
} from "@vue/reactivity"

const useGroupSubscription = () => {
    const error = ref(null)
    const invite = async (nickmail, groupId) => {
        error.value = null
        let body = {
            nickmail: nickmail,
            group_id: groupId
        }
        try {
            let res = await fetch('/api/group/invite', {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(body)
            })
            if (!res.ok) {
                if (res.status == 404) {
                    throw Error("could not find such user")
                } 

                if (res.status == 400) {
                    throw Error("you can't invite owner of the group")
                }
            }
        } catch (e) {
            console.log(e.message)
            error.value = e.message
        }
    }
    return {
        error,
        invite
    }
}
export default useGroupSubscription