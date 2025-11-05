import {
  Card,
  CardContent,
  CardFooter,
} from '@/components/ui/card';
import { Metadata } from 'next';
import Image from 'next/image';
import CredentialsSignUpForm from './credentials-signup-form';

export const metadata: Metadata = {
  title: 'Register',
};


const SignInPage = async () => {
  return (
    <>
      <div className="w-full max-w-md flex justify-center items-center flex-col mb-10">
        <Image src="/logo.svg" alt="Raven" width={150} height={150} className="w-auto h-auto pb-3" priority={true} />
      </div>
      <div className="w-full max-w-lg">
        <h1 className="font-semibold text-3xl text-center mb-5">Create your <span className="text-purple-500">Resturant</span> account</h1>
        <Card className="w-full max-w-lg px-4 py-8">
          <CardContent>
            <CredentialsSignUpForm />
          </CardContent>
          <CardFooter className="justify-center text-sm text-muted-foreground">
            Already have an account? <a className="underline ml-2" href="/auth/login">Sign In</a>
          </CardFooter>
        </Card>
      </div >
    </>
  )
};

export default SignInPage;
