import { Navigate } from "react-router-dom";

interface ProtectedRouteProps {
  token: string;
  redirectPath: string;
  children: JSX.Element;
}

export default function ProtectedRoute({
  token,
  redirectPath,
  children,
}: ProtectedRouteProps): JSX.Element {
  if (!token) {
    return <Navigate to={redirectPath} replace />;
  }

  return children;
}
