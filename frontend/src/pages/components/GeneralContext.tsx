import { Outlet } from "react-router-dom"
import { useAppSelector } from "../../store/hooks"

 const GeneralContext = () => {

    const token = useAppSelector((store) => store.userSlice.accessToken)
    return (
        <div>
        <Outlet context={{token}}/>
        </div>
    )
}

export default GeneralContext
