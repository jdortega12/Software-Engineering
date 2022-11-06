import UserProfileScreen from "./UserProfileScreen"
import UserProfileScreenPersonal from "./UserProfileScreenPersonal"
import UserProfileScreenNotPersonal from "./UserProfileScreenNotPersonal"

test("User Profile Render Smoke Test", () => {
    <UserProfileScreen username="bingus" isSelf={true}/>
})

test("Personal Profile Render Smoke Test", () => {
    eh = UserProfileScreenPersonal("test")
})

test("Generic Profile Render Smoke Test", () => {
    meh = UserProfileScreenNotPersonal("test")
})