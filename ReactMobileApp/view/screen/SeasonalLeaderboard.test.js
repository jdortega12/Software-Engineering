import SeasonalLeaderboard from "./SeasonalLeaderboard"
import TestRenderer from "react-test-renderer"

test("create account screen smoke test ", ()=> {
    const renderer = TestRenderer.create(<SeasonalLeaderboard/>)
    const instance = renderer.root
})