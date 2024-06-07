"use client";
import { login } from "@/services/auth.service";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

const formSchema = z.object({
  email: z.string().email(),
  password: z
    .string()
    .min(2, { message: "Password must be at least of 6 characters." }),
});

const useContainer = () => {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    const res = await login(values);
    console.log(res);
  }

  return {
    // functions
    onSubmit,
    // state
    form,
  };
};

export default useContainer;
