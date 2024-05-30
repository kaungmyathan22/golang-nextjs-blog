import { AppRoutes } from "./constants";
type Routes = {
  [key: string]: Breadcrumb[];
};

const postsBreadCrumb = [
  {
    name: "Dashboard",
    path: AppRoutes.DASHBOARD,
  },
  {
    name: "Posts",
    path: AppRoutes.POSTS,
  },
];
const postDetailsRegex = /\/posts\/1(.*)/;

export function getBreadcumb(path: string): Array<Breadcrumb> {
  const routes: Routes = {
    [AppRoutes.DASHBOARD]: [
      {
        name: "Dashboard",
      },
    ],
    [AppRoutes.POSTS]: [
      ...postsBreadCrumb,
      {
        name: "All Post",
      },
    ],
    [AppRoutes.CREATE_POSTS]: [
      ...postsBreadCrumb,
      {
        name: "Create Post",
      },
    ],
  };
  if (path.startsWith("/posts")) {
    if (postDetailsRegex.test(path)) {
      return [...postsBreadCrumb, { name: "Edit Post" }];
    }
  }
  return routes[path] ?? [];
}
