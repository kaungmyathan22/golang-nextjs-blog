"use client";
import Link from "next/link";

import Client from "@/components/shared/Client";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

export default function ForgotPassword() {
  return (
    <Client>
      <div className="h-screen w-screen grid place-content-center">
        <Card className="mx-auto max-w-sm">
          <CardHeader>
            <CardTitle className="text-2xl">Forgot Password</CardTitle>
            <CardDescription>
              Enter your email below to reset your account
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className="grid gap-4">
              <div className="grid gap-2">
                <Label htmlFor="email">Email</Label>
                <Input
                  id="email"
                  type="email"
                  placeholder="m@example.com"
                  required
                />
              </div>
              <Button type="submit" className="w-full">
                Submit
              </Button>
            </div>
            <div className="mt-4 text-center text-sm">
              Remember your account?{" "}
              <Link href="/signup" className="underline">
                Signin
              </Link>
            </div>
          </CardContent>
        </Card>
      </div>
    </Client>
  );
}
