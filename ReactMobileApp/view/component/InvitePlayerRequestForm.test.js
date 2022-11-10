import TestRenderer from "react-test-renderer"
import InvitePlayerRequestForm from "./InvitePlayerRequestForm"

test("invite player form smoke test ", () => {
    const renderer = TestRenderer.create(<InvitePlayerRequestForm/>)
    const instance = renderer.root
})