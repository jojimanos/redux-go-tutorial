import {
  createBrowserRouter,
  createRoutesFromElements,
  NavLink,
  Route,
  RouteMatch,
  RouteObject,
} from "react-router-dom";
import Login from "../pages/Login";
import LandingPage from "../pages/LandingPage";
import SignupPage from "../pages/Signup";
import OrderingPage from "../pages/OrderingPage";
import ProfilePage from "../pages/ProfilePage";

const routesConfig = createRoutesFromElements(
  <Route>
    <Route path="/" element={<LandingPage />} />
    <Route path="/login" element={<Login />} />
    <Route path="/signup" element={<SignupPage />} />
    <Route path="/profile" element={<ProfilePage />} />
    <Route path="/order" element={<OrderingPage />} />
  </Route>
);

export function setupBrowserRouter(routesConfiguration: RouteObject[]) {
  return createBrowserRouter(routesConfiguration);
}

export default routesConfig;
