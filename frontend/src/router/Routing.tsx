import {
  createBrowserRouter,
  createRoutesFromElements,
  Route,
  RouteObject,
} from "react-router-dom";
import Login from "../pages/Login";
import LandingPage from "../pages/LandingPage";
import SignupPage from "../pages/Signup";
import OrderingPage from "../pages/OrderingPage";
import ProfilePage from "../pages/ProfilePage";
import ProtectedRoute from "./ProtectedRoute";
import GeneralContext from "../pages/components/GeneralContext";

const authToken = window.localStorage.getItem("token");

const routesConfig = createRoutesFromElements(
  <Route element={<GeneralContext/>}>
    <Route path="/" element={<LandingPage />} />
    <Route path="/login" element={<Login />} />
    <Route path="/signup" element={<SignupPage />} />
    <Route
      path="/profile"
      element={
        <ProtectedRoute redirectPath="/login" token={authToken as string}>
          <ProfilePage />
        </ProtectedRoute>
      }
    />
    <Route
      path="/order"
      element={
        <ProtectedRoute redirectPath="/login" token={authToken as string}>
          <OrderingPage />
        </ProtectedRoute>
      }
    />
  </Route>
);

export function setupBrowserRouter(routesConfiguration: RouteObject[]) {
  return createBrowserRouter(routesConfiguration);
}

export default routesConfig;
