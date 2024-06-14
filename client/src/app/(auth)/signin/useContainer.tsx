"use client";
import { useToast } from "@/components/ui/use-toast";
import { State, login } from "@/services/auth.service";
import { zodResolver } from "@hookform/resolvers/zod";
import { useEffect } from "react";
import { useFormState } from "react-dom";
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

  const [state, formAction] = useFormState<State, FormData>(login, null);

  const { toast } = useToast();
  useEffect(() => {
    toast({
      title: state?.status === "success" ? "Success" : "Error",
      description: state?.message,
      variant: state?.status === "success" ? "success" : "destructive",
    });
  }, [state, toast]);

  return {
    // functions
    formAction,
    // state
    form,
  };
};

export default useContainer;
