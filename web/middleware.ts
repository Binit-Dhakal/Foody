import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

// Routes that should only be accessed by guests (not logged-in users)
const guestOnlyRoutes = ['/auth/login', '/auth/register', '/auth/vendor-register'];

export function middleware(req: NextRequest) {
  const url = req.nextUrl.clone();
  const token = req.cookies.get('accessToken')?.value;

  // If logged-in user tries to access guest-only page → redirect
  if (token && guestOnlyRoutes.some(route => url.pathname.startsWith(route))) {
    url.pathname = '/'; // redirect to homepage; role-specific redirect handled client-side
    return NextResponse.redirect(url);
  }

  return NextResponse.next();
}

export const config = {
  matcher: ['/auth/:path*'], // only run middleware on auth pages
};

