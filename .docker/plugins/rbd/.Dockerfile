FROM centos:7.3.1611

ENV CEPH_VERSION kraken

RUN rpm --import 'https://download.ceph.com/keys/release.asc'
RUN rpm -Uvh http://download.ceph.com/rpm-${CEPH_VERSION}/el7/noarch/ceph-release-1-1.el7.noarch.rpm
RUN yum install -y epel-release && yum clean all
RUN yum install -y ceph-common e2fsprogs xfsprogs iproute && yum clean all

RUN mkdir -p /etc/rexray /run/docker/plugins /var/lib/libstorage/volumes
ADD rexray /usr/bin/rexray
ADD rexray.yml /etc/rexray/rexray.yml

ADD .rexray.sh /rexray.sh
RUN chmod +x /rexray.sh

CMD [ "rexray", "start", "-f", "--nopid" ]
ENTRYPOINT [ "/rexray.sh" ]
