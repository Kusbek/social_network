import { ref } from "@vue/reactivity"

const useGroupSubscription = () => {
    const error = ref(null)
    const invite = async (nickmail) => {
        error.value = `User with username or email ${nickmail} was not found`
    }
    return {
        error,
        invite
    }
}
export default useGroupSubscription