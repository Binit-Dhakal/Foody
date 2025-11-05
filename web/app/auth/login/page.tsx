import {
    Card,
    CardContent,
    CardFooter,
} from '@/components/ui/card';
import { Metadata } from 'next';
import Image from 'next/image';
import CredentialsSignInForm from './credentials-signin-form';

export const metadata: Metadata = {
    title: 'Sign In',
};

const SignInPage = async () => {
    return (
        <>
            <div className="w-full max-w-md flex justify-center items-center flex-col mb-10">
                <Image src="/logo.svg" alt="Raven" width={150} height={150} className="w-auto h-auto pb-3" priority={true} />
            </div>
            <div className="w-full max-w-md">
                <h1 className="font-semibold text-3xl text-center mb-5">Sign in to your <span className="text-purple-500">Foody</span> account</h1>
                <Card className="w-full max-w-md px-4 py-8">
                    <CardContent>
                        <CredentialsSignInForm />
                    </CardContent>
                    <CardFooter className="justify-center text-sm text-muted-foreground">
                        Don't have an account? <a className="underline ml-2" href="/auth/register">Sign up</a>
                    </CardFooter>
                </Card>
            </div >
        </>
    )
};

export default SignInPage;
