import LoginScreen from "../LoginScreen"
import TestRenderer from "react-test-renderer"

test("login screen smoke test", () => {
    const renderer = TestRenderer.create(<LoginScreen/>)
    const instance = renderer.root
})