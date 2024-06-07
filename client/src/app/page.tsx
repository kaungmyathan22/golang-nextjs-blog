"use client";

import { Button } from "@/components/ui/button";
import Link from "next/link";

export default function Home() {
  return (
    <div className="w-full lg:grid lg:min-h-screen lg:grid-cols-2">
      <div className="flex items-center justify-center">
        <div className="mx-auto grid w-[350px] gap-6">
          <div className="grid gap-2 text-center">
            <h1 className="text-3xl font-bold">Welcome</h1>
          </div>
          <div className="flex items-center gap-4">
            <Link href={"/signin"} className="w-full">
              <Button type="submit" className="w-full">
                Login
              </Button>
            </Link>
            <Link href={"/signup"} className="w-full">
              <Button className="w-full">Register</Button>
            </Link>
          </div>
        </div>
      </div>
      <div className="hidden bg-muted lg:block bg-black"></div>
    </div>
  );
}
