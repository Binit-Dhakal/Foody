import "@/app/globals.css";
import Image from "next/image";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div className="w-full min-h-screen flex items-center justify-center flex-row">
      <div className="w-full max-w-md flex justify-center items-center flex-col mb-10">
        <Image src="/logo.svg" alt="Raven" width={150} height={150} className="w-auto h-auto pb-3" priority={true} />
      </div>
      {children}
    </div>
  );
}
