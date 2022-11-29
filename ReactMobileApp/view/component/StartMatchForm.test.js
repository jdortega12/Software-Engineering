import TestRenderer from "react-test-renderer"
import StartMatchForm from "./StartMatchForm"

test("StartMatchForm smoke test ", () => {
    const renderer = TestRenderer.create(<StartMatchForm/>)
    const instance = renderer.root
})