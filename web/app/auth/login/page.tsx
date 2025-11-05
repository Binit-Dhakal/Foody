import {
    Card,
    CardContent,
    CardFooter,
} from '@/components/ui/card';
import { Metadata } from 'next';
import CredentialsSignInForm from './credentials-signin-form';

export const metadata: Metadata = {
    title: 'Sign In',
};

const SignInPage = async () => {
    return (
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
    )
};

export default SignInPage;
