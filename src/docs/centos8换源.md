# CentOS 8 换源，设置dnf / yum镜像
    aliyun更新了centos8的说明
    
    1、备份
    mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.backup
    
    2、下载新的CentOS-Base.repo 到/etc/yum.repos.d/
    curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-8.repo
    
    3、生成缓存
    yum makecache
    
    
    4 配置镜像的。
    
    cd /etc/yum.repos.d
    #备份
    cp CentOS-Base.repo CentOS-Base.repo.bak;
    cp CentOS-AppStream.repo CentOS-AppStream.repo.bak;
    cp CentOS-Extras.repo CentOS-Extras.repo.bak;
    
    sed -i 's/mirrorlist=/#mirrorlist=/g' CentOS-Base.repo CentOS-AppStream.repo CentOS-Extras.repo;
    sed -i 's/#baseurl=/baseurl=/g' CentOS-Base.repo CentOS-AppStream.repo CentOS-Extras.repo;
    sed -i 's/http:\/\/mirror.centos.org/https:\/\/mirrors.aliyun.com/g' CentOS-Base.repo CentOS-AppStream.repo CentOS-Extras.repo
    
# ssh 生成密钥
    ssh-keygen -q -t rsa -b 2048 -f /etc/ssh/ssh_host_rsa_key -N '' ; 
    ssh-keygen -q -t ecdsa -f /etc/ssh/ssh_host_ecdsa_key -N '';
    ssh-keygen -t dsa -f /etc/ssh/ssh_host_ed25519_key -N '';
    
    8、配置ssh无密码登录
    先退出 ssh客户端
    ssh-keygen -t rsa
    cd ~/.ssh
    cat id_rsa.pub >> authorized_keys
    
# 安装 wxWidgets
    先安装 
        安装wxWidgets相关依赖，不装这个就会报以下错误
        yum -y install ncurses-devel unixODBC-devel openssl-devel  gcc gcc-c++ autoconf automake
        yum -y install gtk2-devel binutils-devel
          
    编译安装 
    ./configure --with-regex=builtin --with-gtk --enable-unicode --disable-shared --prefix=/usr/local/wxWidgets
    make && make install  
    sudo ldconfig      
    
# erlang 简单安装 
    wxWidgets插件
    1- 添加源
          # wget http://packages.erlang-solutions.com/erlang-solutions-1.0-1.noarch.rpm
          # rpm -Uvh erlang-solutions-1.0-1.noarch.rpm
          # rpm --import http://packages.erlang-solutions.com/rpm/erlang_solutions.asc
    
    2- 看一下新装上的源
          # vim /etc/yum.repos.d/erlang_solutions.repo
                [erlang-solutions]
                name=Centos $releasever - $basearch - Erlang Solutions
                baseurl=http://packages.erlang-solutions.com/rpm/centos/$releasever/$basearch
                gpgcheck=0
                gpgkey=http://packages.erlang-solutions.com/debian/erlang_solutions.asc
                enabled=1
    
    3- 安装
          # yum install erlang erlang-wx
    
          按照源的地址在网页上下载
          http://packages.erlang-solutions.com/rpm/centos/7/x86_64/
          erlang-18.0-1.el7.centos.x86_64.rpm 和 erlang-wx-18.0-1.el7.centos.x86_64.rpm
          # yum install erlang-18.0-1.el7.centos.x86_64.rpm erlang-wx-18.0-1.el7.centos.x86_64.rpm    
    
# erlang编译安装
        安装 erlang
        3 安装java环境 
        3.1 下载java安装文件 wget http://download.oracle.com/otn-pub/java/jdk/7u1-b08/jdk-7u1-linux-i586.rpm 
        3.2 使用rpm 安装 rpm -ivh jdk-7u1-linux-i586.rpm 
        yum -y install ncurses-devel unixODBC-devel openssl-devel  gcc gcc-c++ autoconf automake libxslt gtk3-devel.x86_64 
        
        运行configure配置 
        ./otp_build autoconf       
        ./configure --prefix=/usr/local/erlang  --with-ssl  -enable-threads -enable-smmp-support -enable-kernel-poll  --enable-hipe  --without-javac
        
        fop is missing可以忽略
        
        运行make install安装 
        8 安装 
        8.1 将erlang源码包解压到/root目录下（一定要/root目录, 非/root目录需要配置参数，还没搞明白)  
        8.2 运行./configrue 
        8.3 运行 make install 
        
        CentOS添加环境变量
        在Linux CentOS系统上安装完php和MySQL后，为了使用方便，需要将php和mysql命令加到系统命令中，如果在没有添加到环境变量之前，执行“php -v”命令查看当前php版本信息时时，则会提示命令不存在的错误，下面我们详细介绍一下在linux下将php和mysql加入到环境变量中的方法（假设php和mysql分别安装在/usr/local/webserver/php/和/usr/local/webserver/mysql/中）。
        
        方法一：直接运行命令export PATH=$PATH:/usr/local/webserver/php/bin 和 export PATH=$PATH:/usr/local/webserver/mysql/bin
        
        使用这种方法，只会对当前会话有效，也就是说每当登出或注销系统以后，PATH 设置就会失效，只是临时生效。
        
        方法二：执行vi ~/.bash_profile修改文件中PATH一行，将/usr/local/webserver/php/bin 和 /usr/local/webserver/mysql/bin 加入到PATH=$PATH:$HOME/bin一行之后
        
        这种方法只对当前登录用户生效
        
        方法三：修改/etc/profile文件使其永久性生效，并对所有系统用户生效，在文件末尾加上如下两行代码
        PATH=$PATH:/usr/local/webserver/php/bin:/usr/local/webserver/mysql/bin
        PATH=$PATH:/usr/local/erlang/bin

        export PATH
        
        最后：执行 命令source /etc/profile或 执行点命令 ./profile使其修改生效，执行完可通过echo $PATH命令查看是否添加成功。