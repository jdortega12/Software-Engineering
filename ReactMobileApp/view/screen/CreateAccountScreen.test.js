import CreateAccountScreen from "./CreateAccountScreen"
import TestRenderer from "react-test-renderer"

test("create account screen smoke test ", ()=> {
    const renderer = TestRenderer.create(<CreateAccountScreen/>)
    const instance = renderer.root
})