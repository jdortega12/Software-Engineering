import TestRenderer from "react-test-renderer"
import CreateTeamScreen from "./CreateTeam"

test("create team screen smoke test ", () => {
    const renderer = TestRenderer.create(<CreateTeamScreen/>)
    const instance = renderer.root
})