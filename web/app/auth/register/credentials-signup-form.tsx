'use client';

import { Input } from '@/components/ui/input';
import { useSearchParams } from 'next/navigation';
import { useAuth } from '@/context/AuthContext';
import { signUpUser } from '@/lib/api';
import { useRouter } from 'next/navigation';
import { useState } from 'react';
import { Button } from '@/components/ui/button';
import { Field, FieldError, FieldLabel } from '@/components/ui/field';
import { SignUpUserSchema } from '@/lib/validators';

const CredentialsSignUpForm = () => {
  const searchParams = useSearchParams();
  const callbackUrl = searchParams.get('callbackUrl') || '/auth/login';
  const router = useRouter()
  const { reloadSession } = useAuth()

  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');
  const [formErrors, setFormErrors] = useState<any>({});

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    setIsLoading(true)
    setFormErrors({});
    setError('')

    const formData = new FormData(e.currentTarget);
    const formDataObj = {
      email: formData.get('email'),
      password: formData.get('password'),
      confirmPassword: formData.get('confirmPassword'),
      fullName: formData.get("fullName"),
    }

    const result = SignUpUserSchema.safeParse(formDataObj);

    if (!result.success) {
      const errors: any = {};
      result.error?.issues.forEach((err) => {
        errors[err.path[0]] = err.message;
      })
      setFormErrors(errors)
      setIsLoading(false)
      return;
    }

    try {
      await signUpUser(result.data)
      await reloadSession()
      router.push(callbackUrl)
    } catch (err: any) {
      setError(err.response?.data?.message || 'Login failed')
    } finally {
      setIsLoading(false)
    }
  };
  return (
    <form onSubmit={handleSubmit} className="space-y-8">
      <Field>
        <FieldLabel htmlFor='fullName'>Full Name</FieldLabel>
        <Input id="fullName" name="fullName" type="text" required />
        <FieldError> {formErrors.fullName}</FieldError>
      </Field>

      <Field>
        <FieldLabel htmlFor='email'>Email</FieldLabel>
        <Input id="email" name="email" type="email" required />
        <FieldError>{formErrors.email}</FieldError>
      </Field>

      <Field>
        <FieldLabel htmlFor='password'>Password</FieldLabel>
        <Input id="password" name="password" type="password" minLength={8} required />
        <FieldError>{formErrors.password}</FieldError>
      </Field>

      <Field>
        <FieldLabel htmlFor='confirmPassword'>Confirm Password</FieldLabel>
        <Input id="confirmPassword" name="confirmPassword" type="password" minLength={8} required />
        <FieldError>{formErrors.confirmPassword}</FieldError>
      </Field>

      {error && <p className="text-red-500 text-sm">{error}</p>}
      <Button type="submit" className="w-full" disabled={isLoading}>
        {isLoading ? "Registering..." : "Register"}
      </Button>
    </form>
  )
};

export default CredentialsSignUpForm;
