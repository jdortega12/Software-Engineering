import TestRenderer from "react-test-renderer"
import AskManagerRequestForm from "./AskManagerRequestForm"

test("ask manager request form smoke test", () => {
    const renderer = TestRenderer.create(<AskManagerRequestForm/>)
    const instance = renderer.root
})