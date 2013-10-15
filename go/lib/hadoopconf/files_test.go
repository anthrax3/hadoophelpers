package hadoopconf

var coreSite = `
<?xml version="1.0"?>
<?xml-stylesheet type="text/xsl" href="configuration.xsl"?>

<!-- Put site-specific property overrides in this file. -->

<configuration>
     <property>
         <name>custom.property</name>
         <value>custom.value</value>
     </property>
     <property>
         <name>fs.default.name</name>
         <value>hdfs://localhost:8020</value>
     </property>
     <property>
        <name>hadoop.security.authentication</name>
        <value>kerberos</value>
     </property>
    <property>   
        <name>hadoop.rpc.protection</name>   
        <value>authentication</value>   
        <description>This is an [OPTIONAL] setting. If not set, defaults to authentication.authentication= authentication only; the client and server mutually authenticate during connection setup.integrity = authentication and integrity; guarantees the integrity of data exchanged between client and server aswell as authentication.privacy = authentication, integrity, and confidentiality; guarantees that data exchanged between client andserver is encrypted and is not readable by a “man in the middle”.
        </description>
    </property>
    <property>  
        <name>hadoop.security.authorization</name>  
        <value>true</value>  
        <description>Enable authorization for different protocols.  
        </description> 
    </property>  
    <property>
        <name>hadoop.security.auth_to_local</name>    
        <value>RULE:[2:$1@$0]([jt]t@.*EXAMPLE.COM)s/.*/$MAPRED_USER/
RULE:[2:$1@$0]([nd]n@.*EXAMPLE.COM)s/.*/$HDFS_USER/
DEFAULT</value> 
        <description>The mapping from Kerberos principal names to local OS user names. </description>
    </property>

    <property>  
        <name>java.security.krb5.conf</name>  
	<value>/Users/eleibovi/work/akamai/hadoop_rpc/target/krb5.conf</value>  
        <description>Enable authorization for different protocols.</description> 
    </property>  
</configuration>
`

var coreDefault = `
<?xml version="1.0"?>
<?xml-stylesheet type="text/xsl" href="configuration.xsl"?>

<!--
   Licensed to the Apache Software Foundation (ASF) under one or more
   contributor license agreements.  See the NOTICE file distributed with
   this work for additional information regarding copyright ownership.
   The ASF licenses this file to You under the Apache License, Version 2.0
   (the "License"); you may not use this file except in compliance with
   the License.  You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
-->

<!-- Do not modify this file directly.  Instead, copy entries that you -->
<!-- wish to modify from this file into core-site.xml and change them -->
<!-- there.  If core-site.xml does not already exist, create it.      -->

<configuration>

<!--- global properties -->

<property>
  <name>hadoop.common.configuration.version</name>
  <value>0.23.0</value>
  <description>version of this configuration file</description>
</property>

<property>
  <name>hadoop.tmp.dir</name>
  <value>/tmp/hadoop-${user.name}</value>
  <description>A base for other temporary directories.</description>
</property>

<property>
  <name>io.native.lib.available</name>
  <value>true</value>
  <description>Should native hadoop libraries, if present, be used.</description>
</property>

<property>
  <name>hadoop.http.filter.initializers</name>
  <value>org.apache.hadoop.http.lib.StaticUserWebFilter</value>
  <description>A comma separated list of class names. Each class in the list 
  must extend org.apache.hadoop.http.FilterInitializer. The corresponding 
  Filter will be initialized. Then, the Filter will be applied to all user 
  facing jsp and servlet web pages.  The ordering of the list defines the 
  ordering of the filters.</description>
</property>

<!--- security properties -->

<property>
  <name>hadoop.security.authorization</name>
  <value>false</value>
  <description>Is service-level authorization enabled?</description>
</property>

<property>
  <name>hadoop.security.instrumentation.requires.admin</name>
  <value>false</value>
  <description>
    Indicates if administrator ACLs are required to access
    instrumentation servlets (JMX, METRICS, CONF, STACKS).
  </description>
</property>

<property>
  <name>hadoop.security.authentication</name>
  <value>simple</value>
  <description>Possible values are simple (no authentication), and kerberos
  </description>
</property>

<property>
  <name>hadoop.security.group.mapping</name>
  <value>org.apache.hadoop.security.JniBasedUnixGroupsMappingWithFallback</value>
  <description>
    Class for user to group mapping (get groups for a given user) for ACL. 
    The default implementation,
    org.apache.hadoop.security.JniBasedUnixGroupsMappingWithFallback, 
    will determine if the Java Native Interface (JNI) is available. If JNI is 
    available the implementation will use the API within hadoop to resolve a 
    list of groups for a user. If JNI is not available then the shell 
    implementation, ShellBasedUnixGroupsMapping, is used.  This implementation 
    shells out to the Linux/Unix environment with the 
    <code>bash -c groups</code> command to resolve a list of groups for a user.
  </description>
</property>

<property>
  <name>hadoop.security.groups.cache.secs</name>
  <value>300</value>
  <description>
    This is the config controlling the validity of the entries in the cache
    containing the user->group mapping. When this duration has expired,
    then the implementation of the group mapping provider is invoked to get
    the groups of the user and then cached back.
  </description>
</property>

<property>
  <name>hadoop.security.group.mapping.ldap.url</name>
  <value></value>
  <description>
    The URL of the LDAP server to use for resolving user groups when using
    the LdapGroupsMapping user to group mapping.
  </description>
</property>

<property>
  <name>hadoop.security.group.mapping.ldap.ssl</name>
  <value>false</value>
  <description>
    Whether or not to use SSL when connecting to the LDAP server.
  </description>
</property>

<property>
  <name>hadoop.security.group.mapping.ldap.ssl.keystore</name>
  <value></value>
  <description>
    File path to the SSL keystore that contains the SSL certificate required
    by the LDAP server.
  </description>
</property>

<property>
  <name>hadoop.security.group.mapping.ldap.ssl.keystore.password.file</name>
  <value></value>
  <description>
    The path to a file containing the password of the LDAP SSL keystore.

    IMPORTANT: This file should be readable only by the Unix user running
    the daemons.
  </description>
</property>

<property>
  <name>hadoop.security.group.mapping.ldap.bind.user</name>
  <value></value>
  <description>
    The distinguished name of the user to bind as when connecting to the LDAP
    server. This may be left blank if the LDAP server supports anonymous binds.
  </description>
</property>

<property>
  <name>hadoop.security.group.mapping.ldap.bind.password.file</name>
  <value></value>
  <description>
    The path to a file containing the password of the bind user.

    IMPORTANT: This file should be readable only by the Unix user running
    the daemons.
  </description>
</property>

<property>
  <name>hadoop.security.group.mapping.ldap.base</name>
  <value></value>
  <description>
    The search base for the LDAP connection. This is a distinguished name,
    and will typically be the root of the LDAP directory.
  </description>
</property>

<property>
  <name>hadoop.security.group.mapping.ldap.search.filter.user</name>
  <value>(&amp;(objectClass=user)(sAMAccountName={0}))</value>
  <description>
    An additional filter to use when searching for LDAP users. The default will
    usually be appropriate for Active Directory installations. If connecting to
    an LDAP server with a non-AD schema, this should be replaced with
    (&amp;(objectClass=inetOrgPerson)(uid={0}). {0} is a special string used to
    denote where the username fits into the filter.
  </description>
</property>

<property>
  <name>hadoop.security.group.mapping.ldap.search.filter.group</name>
  <value>(objectClass=group)</value>
  <description>
    An additional filter to use when searching for LDAP groups. This should be
    changed when resolving groups against a non-Active Directory installation.
    posixGroups are currently not a supported group class.
  </description>
</property>

<property>
  <name>hadoop.security.group.mapping.ldap.search.attr.member</name>
  <value>member</value>
  <description>
    The attribute of the group object that identifies the users that are
    members of the group. The default will usually be appropriate for
    any LDAP installation.
  </description>
</property>

<property>
  <name>hadoop.security.group.mapping.ldap.search.attr.group.name</name>
  <value>cn</value>
  <description>
    The attribute of the group object that identifies the group name. The
    default will usually be appropriate for all LDAP systems.
  </description>
</property>

<property>
  <name>hadoop.security.group.mapping.ldap.directory.search.timeout</name>
  <value>10000</value>
  <description>
    The attribute applied to the LDAP SearchControl properties to set a
    maximum time limit when searching and awaiting a result.
    Set to 0 if infinite wait period is desired.
    Default is 10 seconds. Units in milliseconds.
  </description>
</property>

<property>
  <name>hadoop.security.service.user.name.key</name>
  <value></value>
  <description>
    For those cases where the same RPC protocol is implemented by multiple
    servers, this configuration is required for specifying the principal
    name to use for the service when the client wishes to make an RPC call.
  </description>
</property>


<property>
    <name>hadoop.security.uid.cache.secs</name>
    <value>14400</value>
    <description>
        This is the config controlling the validity of the entries in the cache
        containing the userId to userName and groupId to groupName used by
        NativeIO getFstat().
    </description>
</property>

<property>
  <name>hadoop.rpc.protection</name>
  <value>authentication</value>
  <description>This field sets the quality of protection for secured sasl 
      connections. Possible values are authentication, integrity and privacy.
      authentication means authentication only and no integrity or privacy; 
      integrity implies authentication and integrity are enabled; and privacy 
      implies all of authentication, integrity and privacy are enabled.
  </description>
</property>

<property>
  <name>hadoop.work.around.non.threadsafe.getpwuid</name>
  <value>false</value>
  <description>Some operating systems or authentication modules are known to
  have broken implementations of getpwuid_r and getpwgid_r, such that these
  calls are not thread-safe. Symptoms of this problem include JVM crashes
  with a stack trace inside these functions. If your system exhibits this
  issue, enable this configuration parameter to include a lock around the
  calls as a workaround.

  An incomplete list of some systems known to have this issue is available
  at http://wiki.apache.org/hadoop/KnownBrokenPwuidImplementations
  </description>
</property>

<property>
  <name>hadoop.kerberos.kinit.command</name>
  <value>kinit</value>
  <description>Used to periodically renew Kerberos credentials when provided
  to Hadoop. The default setting assumes that kinit is in the PATH of users
  running the Hadoop client. Change this to the absolute path to kinit if this
  is not the case.
  </description>
</property>

<property>
  <name>hadoop.security.auth_to_local</name>
  <value></value>
  <description>Maps kerberos principals to local user names</description>
</property>

<!-- i/o properties -->
<property>
  <name>io.file.buffer.size</name>
  <value>4096</value>
  <description>The size of buffer for use in sequence files.
  The size of this buffer should probably be a multiple of hardware
  page size (4096 on Intel x86), and it determines how much data is
  buffered during read and write operations.</description>
</property>
  
<property>
  <name>io.bytes.per.checksum</name>
  <value>512</value>
  <description>The number of bytes per checksum.  Must not be larger than
  io.file.buffer.size.</description>
</property>

<property>
  <name>io.skip.checksum.errors</name>
  <value>false</value>
  <description>If true, when a checksum error is encountered while
  reading a sequence file, entries are skipped, instead of throwing an
  exception.</description>
</property>

<property>
  <name>io.compression.codecs</name>
  <value></value>
  <description>A comma-separated list of the compression codec classes that can
  be used for compression/decompression. In addition to any classes specified
  with this property (which take precedence), codec classes on the classpath
  are discovered using a Java ServiceLoader.</description>
</property>

<property>
  <name>io.compression.codec.bzip2.library</name>
  <value>system-native</value>
  <description>The native-code library to be used for compression and
  decompression by the bzip2 codec.  This library could be specified
  either by by name or the full pathname.  In the former case, the
  library is located by the dynamic linker, usually searching the
  directories specified in the environment variable LD_LIBRARY_PATH.
  
  The value of "system-native" indicates that the default system
  library should be used.  To indicate that the algorithm should
  operate entirely in Java, specify "java-builtin".</description>
</property>

<property>
  <name>io.serializations</name>
  <value>org.apache.hadoop.io.serializer.WritableSerialization,org.apache.hadoop.io.serializer.avro.AvroSpecificSerialization,org.apache.hadoop.io.serializer.avro.AvroReflectSerialization</value>
  <description>A list of serialization classes that can be used for
  obtaining serializers and deserializers.</description>
</property>

<property>
  <name>io.seqfile.local.dir</name>
  <value>${hadoop.tmp.dir}/io/local</value>
  <description>The local directory where sequence file stores intermediate
  data files during merge.  May be a comma-separated list of
  directories on different devices in order to spread disk i/o.
  Directories that do not exist are ignored.
  </description>
</property>

<property>
  <name>io.map.index.skip</name>
  <value>0</value>
  <description>Number of index entries to skip between each entry.
  Zero by default. Setting this to values larger than zero can
  facilitate opening large MapFiles using less memory.</description>
</property>

<property>
  <name>io.map.index.interval</name>
  <value>128</value>
  <description>
    MapFile consist of two files - data file (tuples) and index file
    (keys). For every io.map.index.interval records written in the
    data file, an entry (record-key, data-file-position) is written
    in the index file. This is to allow for doing binary search later
    within the index file to look up records by their keys and get their
    closest positions in the data file.
  </description>
</property>

<!-- file system properties -->

<property>
  <name>fs.defaultFS</name>
  <value>file:///</value>
  <description>The name of the default file system.  A URI whose
  scheme and authority determine the FileSystem implementation.  The
  uri's scheme determines the config property (fs.SCHEME.impl) naming
  the FileSystem implementation class.  The uri's authority is used to
  determine the host, port, etc. for a filesystem.</description>
</property>

<property>
  <name>fs.default.name</name>
  <value>file:///</value>
  <description>Deprecated. Use (fs.defaultFS) property
  instead</description>
</property>

<property>
  <name>fs.trash.interval</name>
  <value>0</value>
  <description>Number of minutes after which the checkpoint
  gets deleted.  If zero, the trash feature is disabled.
  This option may be configured both on the server and the
  client. If trash is disabled server side then the client
  side configuration is checked. If trash is enabled on the
  server side then the value configured on the server is
  used and the client configuration value is ignored.
  </description>
</property>

<property>
  <name>fs.trash.checkpoint.interval</name>
  <value>0</value>
  <description>Number of minutes between trash checkpoints.
  Should be smaller or equal to fs.trash.interval. If zero,
  the value is set to the value of fs.trash.interval.
  Every time the checkpointer runs it creates a new checkpoint 
  out of current and removes checkpoints created more than 
  fs.trash.interval minutes ago.
  </description>
</property>

<property>
  <name>fs.AbstractFileSystem.file.impl</name>
  <value>org.apache.hadoop.fs.local.LocalFs</value>
  <description>The AbstractFileSystem for file: uris.</description>
</property>


<property>
  <name>fs.AbstractFileSystem.hdfs.impl</name>
  <value>org.apache.hadoop.fs.Hdfs</value>
  <description>The FileSystem for hdfs: uris.</description>
</property>

<property>
  <name>fs.AbstractFileSystem.viewfs.impl</name>
  <value>org.apache.hadoop.fs.viewfs.ViewFs</value>
  <description>The AbstractFileSystem for view file system for viewfs: uris
  (ie client side mount table:).</description>
</property>

<property>
  <name>fs.ftp.host</name>
  <value>0.0.0.0</value>
  <description>FTP filesystem connects to this server</description>
</property>

<property>
  <name>fs.ftp.host.port</name>
  <value>21</value>
  <description>
    FTP filesystem connects to fs.ftp.host on this port
  </description>
</property>

<property>
  <name>fs.df.interval</name>
  <value>60000</value>
  <description>Disk usage statistics refresh interval in msec.</description>
</property>

<property>
  <name>fs.s3.block.size</name>
  <value>67108864</value>
  <description>Block size to use when writing files to S3.</description>
</property>

<property>
  <name>fs.s3.buffer.dir</name>
  <value>${hadoop.tmp.dir}/s3</value>
  <description>Determines where on the local filesystem the S3 filesystem
  should store files before sending them to S3
  (or after retrieving them from S3).
  </description>
</property>

<property>
  <name>fs.s3.maxRetries</name>
  <value>4</value>
  <description>The maximum number of retries for reading or writing files to S3, 
  before we signal failure to the application.
  </description>
</property>

<property>
  <name>fs.s3.sleepTimeSeconds</name>
  <value>10</value>
  <description>The number of seconds to sleep between each S3 retry.
  </description>
</property>


<property>
  <name>fs.automatic.close</name>
  <value>true</value>
  <description>By default, FileSystem instances are automatically closed at program
  exit using a JVM shutdown hook. Setting this property to false disables this
  behavior. This is an advanced option that should only be used by server applications
  requiring a more carefully orchestrated shutdown sequence.
  </description>
</property>

<property>
  <name>fs.s3n.block.size</name>
  <value>67108864</value>
  <description>Block size to use when reading files using the native S3
  filesystem (s3n: URIs).</description>
</property>

<property>
  <name>io.seqfile.compress.blocksize</name>
  <value>1000000</value>
  <description>The minimum block size for compression in block compressed 
          SequenceFiles.
  </description>
</property>

<property>
  <name>io.seqfile.lazydecompress</name>
  <value>true</value>
  <description>Should values of block-compressed SequenceFiles be decompressed
          only when necessary.
  </description>
</property>

<property>
  <name>io.seqfile.sorter.recordlimit</name>
  <value>1000000</value>
  <description>The limit on number of records to be kept in memory in a spill 
          in SequenceFiles.Sorter
  </description>
</property>

 <property>
  <name>io.mapfile.bloom.size</name>
  <value>1048576</value>
  <description>The size of BloomFilter-s used in BloomMapFile. Each time this many
  keys is appended the next BloomFilter will be created (inside a DynamicBloomFilter).
  Larger values minimize the number of filters, which slightly increases the performance,
  but may waste too much space if the total number of keys is usually much smaller
  than this number.
  </description>
</property>

<property>
  <name>io.mapfile.bloom.error.rate</name>
  <value>0.005</value>
  <description>The rate of false positives in BloomFilter-s used in BloomMapFile.
  As this value decreases, the size of BloomFilter-s increases exponentially. This
  value is the probability of encountering false positives (default is 0.5%).
  </description>
</property>

<property>
  <name>hadoop.util.hash.type</name>
  <value>murmur</value>
  <description>The default implementation of Hash. Currently this can take one of the
  two values: 'murmur' to select MurmurHash and 'jenkins' to select JenkinsHash.
  </description>
</property>


<!-- ipc properties -->

<property>
  <name>ipc.client.idlethreshold</name>
  <value>4000</value>
  <description>Defines the threshold number of connections after which
               connections will be inspected for idleness.
  </description>
</property>

<property>
  <name>ipc.client.kill.max</name>
  <value>10</value>
  <description>Defines the maximum number of clients to disconnect in one go.
  </description>
</property>

<property>
  <name>ipc.client.connection.maxidletime</name>
  <value>10000</value>
  <description>The maximum time in msec after which a client will bring down the
               connection to the server.
  </description>
</property>

<property>
  <name>ipc.client.connect.max.retries</name>
  <value>10</value>
  <description>Indicates the number of retries a client will make to establish
               a server connection.
  </description>
</property>

<property>
  <name>ipc.client.connect.timeout</name>
  <value>20000</value>
  <description>Indicates the number of milliseconds a client will wait for the 
               socket to establish a server connection.
  </description>
</property>

<property>
  <name>ipc.client.connect.max.retries.on.timeouts</name>
  <value>45</value>
  <description>Indicates the number of retries a client will make on socket timeout
               to establish a server connection.
  </description>
</property>

<property>
  <name>ipc.server.listen.queue.size</name>
  <value>128</value>
  <description>Indicates the length of the listen queue for servers accepting
               client connections.
  </description>
</property>

<property>
  <name>ipc.server.tcpnodelay</name>
  <value>false</value>
  <description>Turn on/off Nagle's algorithm for the TCP socket connection on 
  the server. Setting to true disables the algorithm and may decrease latency
  with a cost of more/smaller packets. 
  </description>
</property>

<property>
  <name>ipc.client.tcpnodelay</name>
  <value>false</value>
  <description>Turn on/off Nagle's algorithm for the TCP socket connection on 
  the client. Setting to true disables the algorithm and may decrease latency
  with a cost of more/smaller packets. 
  </description>
</property>


<!-- Proxy Configuration -->

<property>
  <name>hadoop.rpc.socket.factory.class.default</name>
  <value>org.apache.hadoop.net.StandardSocketFactory</value>
  <description> Default SocketFactory to use. This parameter is expected to be
    formatted as "package.FactoryClassName".
  </description>
</property>

<property>
  <name>hadoop.rpc.socket.factory.class.ClientProtocol</name>
  <value></value>
  <description> SocketFactory to use to connect to a DFS. If null or empty, use
    hadoop.rpc.socket.class.default. This socket factory is also used by
    DFSClient to create sockets to DataNodes.
  </description>
</property>



<property>
  <name>hadoop.socks.server</name>
  <value></value>
  <description> Address (host:port) of the SOCKS server to be used by the
    SocksSocketFactory.
  </description>
</property>

<!-- Topology Configuration -->
<property>
  <name>net.topology.node.switch.mapping.impl</name>
  <value>org.apache.hadoop.net.ScriptBasedMapping</value>
  <description> The default implementation of the DNSToSwitchMapping. It
    invokes a script specified in net.topology.script.file.name to resolve
    node names. If the value for net.topology.script.file.name is not set, the
    default value of DEFAULT_RACK is returned for all node names.
  </description>
</property>

<property>
  <name>net.topology.impl</name>
  <value>org.apache.hadoop.net.NetworkTopology</value>
  <description> The default implementation of NetworkTopology which is classic three layer one.
  </description>
</property>

<property>
  <name>net.topology.script.file.name</name>
  <value></value>
  <description> The script name that should be invoked to resolve DNS names to
    NetworkTopology names. Example: the script would take host.foo.bar as an
    argument, and return /rack1 as the output.
  </description>
</property>

<property>
  <name>net.topology.script.number.args</name>
  <value>100</value>
  <description> The max number of args that the script configured with 
    net.topology.script.file.name should be run with. Each arg is an
    IP address.
  </description>
</property>

<property>
  <name>net.topology.table.file.name</name>
  <value></value>
  <description> The file name for a topology file, which is used when the
    net.topology.node.switch.mapping.impl property is set to
    org.apache.hadoop.net.TableMapping. The file format is a two column text
    file, with columns separated by whitespace. The first column is a DNS or
    IP address and the second column specifies the rack where the address maps.
    If no entry corresponding to a host in the cluster is found, then 
    /default-rack is assumed.
  </description>
</property>

<!-- Local file system -->
<property>
  <name>file.stream-buffer-size</name>
  <value>4096</value>
  <description>The size of buffer to stream files.
  The size of this buffer should probably be a multiple of hardware
  page size (4096 on Intel x86), and it determines how much data is
  buffered during read and write operations.</description>
</property>

<property>
  <name>file.bytes-per-checksum</name>
  <value>512</value>
  <description>The number of bytes per checksum.  Must not be larger than
  file.stream-buffer-size</description>
</property>

<property>
  <name>file.client-write-packet-size</name>
  <value>65536</value>
  <description>Packet size for clients to write</description>
</property>

<property>
  <name>file.blocksize</name>
  <value>67108864</value>
  <description>Block size</description>
</property>

<property>
  <name>file.replication</name>
  <value>1</value>
  <description>Replication factor</description>
</property>

<!-- s3 File System -->

<property>
  <name>s3.stream-buffer-size</name>
  <value>4096</value>
  <description>The size of buffer to stream files.
  The size of this buffer should probably be a multiple of hardware
  page size (4096 on Intel x86), and it determines how much data is
  buffered during read and write operations.</description>
</property>

<property>
  <name>s3.bytes-per-checksum</name>
  <value>512</value>
  <description>The number of bytes per checksum.  Must not be larger than
  s3.stream-buffer-size</description>
</property>

<property>
  <name>s3.client-write-packet-size</name>
  <value>65536</value>
  <description>Packet size for clients to write</description>
</property>

<property>
  <name>s3.blocksize</name>
  <value>67108864</value>
  <description>Block size</description>
</property>

<property>
  <name>s3.replication</name>
  <value>3</value>
  <description>Replication factor</description>
</property>

<!-- s3native File System -->

<property>
  <name>s3native.stream-buffer-size</name>
  <value>4096</value>
  <description>The size of buffer to stream files.
  The size of this buffer should probably be a multiple of hardware
  page size (4096 on Intel x86), and it determines how much data is
  buffered during read and write operations.</description>
</property>

<property>
  <name>s3native.bytes-per-checksum</name>
  <value>512</value>
  <description>The number of bytes per checksum.  Must not be larger than
  s3native.stream-buffer-size</description>
</property>

<property>
  <name>s3native.client-write-packet-size</name>
  <value>65536</value>
  <description>Packet size for clients to write</description>
</property>

<property>
  <name>s3native.blocksize</name>
  <value>67108864</value>
  <description>Block size</description>
</property>

<property>
  <name>s3native.replication</name>
  <value>3</value>
  <description>Replication factor</description>
</property>

<!-- FTP file system -->
<property>
  <name>ftp.stream-buffer-size</name>
  <value>4096</value>
  <description>The size of buffer to stream files.
  The size of this buffer should probably be a multiple of hardware
  page size (4096 on Intel x86), and it determines how much data is
  buffered during read and write operations.</description>
</property>

<property>
  <name>ftp.bytes-per-checksum</name>
  <value>512</value>
  <description>The number of bytes per checksum.  Must not be larger than
  ftp.stream-buffer-size</description>
</property>

<property>
  <name>ftp.client-write-packet-size</name>
  <value>65536</value>
  <description>Packet size for clients to write</description>
</property>

<property>
  <name>ftp.blocksize</name>
  <value>67108864</value>
  <description>Block size</description>
</property>

<property>
  <name>ftp.replication</name>
  <value>3</value>
  <description>Replication factor</description>
</property>

<!-- Tfile -->

<property>
  <name>tfile.io.chunk.size</name>
  <value>1048576</value>
  <description>
    Value chunk size in bytes. Default  to
    1MB. Values of the length less than the chunk size is
    guaranteed to have known value length in read time (See also
    TFile.Reader.Scanner.Entry.isValueLengthKnown()).
  </description>
</property>

<property>
  <name>tfile.fs.output.buffer.size</name>
  <value>262144</value>
  <description>
    Buffer size used for FSDataOutputStream in bytes.
  </description>
</property>

<property>
  <name>tfile.fs.input.buffer.size</name>
  <value>262144</value>
  <description>
    Buffer size used for FSDataInputStream in bytes.
  </description>
</property>

<!-- HTTP web-consoles Authentication -->

<property>
  <name>hadoop.http.authentication.type</name>
  <value>simple</value>
  <description>
    Defines authentication used for Oozie HTTP endpoint.
    Supported values are: simple | kerberos | #AUTHENTICATION_HANDLER_CLASSNAME#
  </description>
</property>

<property>
  <name>hadoop.http.authentication.token.validity</name>
  <value>36000</value>
  <description>
    Indicates how long (in seconds) an authentication token is valid before it has
    to be renewed.
  </description>
</property>

<property>
  <name>hadoop.http.authentication.signature.secret.file</name>
  <value>${user.home}/hadoop-http-auth-signature-secret</value>
  <description>
    The signature secret for signing the authentication tokens.
    The same secret should be used for JT/NN/DN/TT configurations.
  </description>
</property>

<property>
  <name>hadoop.http.authentication.cookie.domain</name>
  <value></value>
  <description>
    The domain to use for the HTTP cookie that stores the authentication token.
    In order to authentiation to work correctly across all Hadoop nodes web-consoles
    the domain must be correctly set.
    IMPORTANT: when using IP addresses, browsers ignore cookies with domain settings.
    For this setting to work properly all nodes in the cluster must be configured
    to generate URLs with hostname.domain names on it.
  </description>
</property>

<property>
  <name>hadoop.http.authentication.simple.anonymous.allowed</name>
  <value>true</value>
  <description>
    Indicates if anonymous requests are allowed when using 'simple' authentication.
  </description>
</property>

<property>
  <name>hadoop.http.authentication.kerberos.principal</name>
  <value>HTTP/_HOST@LOCALHOST</value>
  <description>
    Indicates the Kerberos principal to be used for HTTP endpoint.
    The principal MUST start with 'HTTP/' as per Kerberos HTTP SPNEGO specification.
  </description>
</property>

<property>
  <name>hadoop.http.authentication.kerberos.keytab</name>
  <value>${user.home}/hadoop.keytab</value>
  <description>
    Location of the keytab file with the credentials for the principal.
    Referring to the same keytab file Oozie uses for its Kerberos credentials for Hadoop.
  </description>
</property>

<property>
  <name>dfs.ha.fencing.methods</name>
  <value></value>
  <description>
    List of fencing methods to use for service fencing. May contain
    builtin methods (eg shell and sshfence) or user-defined method.
  </description>
</property>

<property>
  <name>dfs.ha.fencing.ssh.connect-timeout</name>
  <value>30000</value>
  <description>
    SSH connection timeout, in milliseconds, to use with the builtin
    sshfence fencer.
  </description>
</property>

<property>
  <name>dfs.ha.fencing.ssh.private-key-files</name>
  <value></value>
  <description>
    The SSH private key files to use with the builtin sshfence fencer.
  </description>
</property>


<!-- Static Web User Filter properties. -->
<property>
  <description>
    The user name to filter as, on static web filters
    while rendering content. An example use is the HDFS
    web UI (user to be used for browsing files).
  </description>
  <name>hadoop.http.staticuser.user</name>
  <value>dr.who</value>
</property>

<property>
  <name>ha.zookeeper.quorum</name>
  <description>
    A list of ZooKeeper server addresses, separated by commas, that are
    to be used by the ZKFailoverController in automatic failover.
  </description>
</property>

<property>
  <name>ha.zookeeper.session-timeout.ms</name>
  <value>5000</value>
  <description>
    The session timeout to use when the ZKFC connects to ZooKeeper.
    Setting this value to a lower value implies that server crashes
    will be detected more quickly, but risks triggering failover too
    aggressively in the case of a transient error or network blip.
  </description>
</property>

<property>
  <name>ha.zookeeper.parent-znode</name>
  <value>/hadoop-ha</value>
  <description>
    The ZooKeeper znode under which the ZK failover controller stores
    its information. Note that the nameservice ID is automatically
    appended to this znode, so it is not normally necessary to
    configure this, even in a federated environment.
  </description>
</property>

<property>
  <name>ha.zookeeper.acl</name>
  <value>world:anyone:rwcda</value>
  <description>
    A comma-separated list of ZooKeeper ACLs to apply to the znodes
    used by automatic failover. These ACLs are specified in the same
    format as used by the ZooKeeper CLI.

    If the ACL itself contains secrets, you may instead specify a
    path to a file, prefixed with the '@' symbol, and the value of
    this configuration will be loaded from within.
  </description>
</property>

<property>
  <name>ha.zookeeper.auth</name>
  <value></value>
  <description>
    A comma-separated list of ZooKeeper authentications to add when
    connecting to ZooKeeper. These are specified in the same format
    as used by the &quot;addauth&quot; command in the ZK CLI. It is
    important that the authentications specified here are sufficient
    to access znodes with the ACL specified in ha.zookeeper.acl.

    If the auths contain secrets, you may instead specify a
    path to a file, prefixed with the '@' symbol, and the value of
    this configuration will be loaded from within.
  </description>
</property>

<!-- SSLFactory configuration -->

<property>
  <name>hadoop.ssl.keystores.factory.class</name>
  <value>org.apache.hadoop.security.ssl.FileBasedKeyStoresFactory</value>
  <description>
    The keystores factory to use for retrieving certificates.
  </description>
</property>

<property>
  <name>hadoop.ssl.require.client.cert</name>
  <value>false</value>
  <description>Whether client certificates are required</description>
</property>

<property>
  <name>hadoop.ssl.hostname.verifier</name>
  <value>DEFAULT</value>
  <description>
    The hostname verifier to provide for HttpsURLConnections.
    Valid values are: DEFAULT, STRICT, STRICT_I6, DEFAULT_AND_LOCALHOST and
    ALLOW_ALL
  </description>
</property>

<property>
  <name>hadoop.ssl.server.conf</name>
  <value>ssl-server.xml</value>
  <description>
    Resource file from which ssl server keystore information will be extracted.
    This file is looked up in the classpath, typically it should be in Hadoop
    conf/ directory.
  </description>
</property>

<property>
  <name>hadoop.ssl.client.conf</name>
  <value>ssl-client.xml</value>
  <description>
    Resource file from which ssl client keystore information will be extracted
    This file is looked up in the classpath, typically it should be in Hadoop
    conf/ directory.
  </description>
</property>

<property>
  <name>hadoop.ssl.enabled</name>
  <value>false</value>
  <description>
    Whether to use SSL for the HTTP endpoints. If set to true, the
    NameNode, DataNode, ResourceManager, NodeManager, HistoryServer and
    MapReduceAppMaster web UIs will be served over HTTPS instead HTTP.
  </description>
</property>

<property>
  <name>hadoop.jetty.logs.serve.aliases</name>
  <value>true</value>
  <description>
    Enable/Disable aliases serving from jetty
  </description>
</property>

<property>
  <name>fs.permissions.umask-mode</name>
  <value>022</value>
  <description>
    The umask used when creating files and directories.
    Can be in octal or in symbolic. Examples are:
    "022" (octal for u=rwx,g=r-x,o=r-x in symbolic),
    or "u=rwx,g=rwx,o=" (symbolic for 007 in octal).
  </description>
</property>

<!-- ha properties -->

<property>
  <name>ha.health-monitor.connect-retry-interval.ms</name>
  <value>1000</value>
  <description>
    How often to retry connecting to the service.
  </description>
</property>

<property>
  <name>ha.health-monitor.check-interval.ms</name>
  <value>1000</value>
  <description>
    How often to check the service.
  </description>
</property>

<property>
  <name>ha.health-monitor.sleep-after-disconnect.ms</name>
  <value>1000</value>
  <description>
    How long to sleep after an unexpected RPC error.
  </description>
</property>

<property>
  <name>ha.health-monitor.rpc-timeout.ms</name>
  <value>45000</value>
  <description>
    Timeout for the actual monitorHealth() calls.
  </description>
</property>

<property>
  <name>ha.failover-controller.new-active.rpc-timeout.ms</name>
  <value>60000</value>
  <description>
    Timeout that the FC waits for the new active to become active
  </description>
</property>

<property>
  <name>ha.failover-controller.graceful-fence.rpc-timeout.ms</name>
  <value>5000</value>
  <description>
    Timeout that the FC waits for the old active to go to standby
  </description>
</property>

<property>
  <name>ha.failover-controller.graceful-fence.connection.retries</name>
  <value>1</value>
  <description>
    FC connection retries for graceful fencing
  </description>
</property>

<property>
  <name>ha.failover-controller.cli-check.rpc-timeout.ms</name>
  <value>20000</value>
  <description>
    Timeout that the CLI (manual) FC waits for monitorHealth, getServiceState
  </description>
</property>

<property>
  <name>ipc.client.fallback-to-simple-auth-allowed</name>
  <value>false</value>
  <description>
    When a client is configured to attempt a secure connection, but attempts to
    connect to an insecure server, that server may instruct the client to
    switch to SASL SIMPLE (unsecure) authentication. This setting controls
    whether or not the client will accept this instruction from the server.
    When false (the default), the client will not allow the fallback to SIMPLE
    authentication, and will abort the connection.
  </description>
</property>

</configuration>
`
