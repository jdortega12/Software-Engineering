import TestRenderer from "react-test-renderer"
import RemovePlayer from "./RemovePlayer"

test("remove player smoke test", () => {
    const renderer = TestRenderer.create(<RemovePlayer/>)
    const instance = renderer.root
})