import TestRenderer from "react-test-renderer"
import AcceptOrDeny from "./AcceptOrDeny"

test("accept or deny smoke test", () => {
    const renderer = TestRenderer.create(<AcceptOrDeny/>)
    const instance = renderer.root
})