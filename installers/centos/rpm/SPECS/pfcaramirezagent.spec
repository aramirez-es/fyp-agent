# Don't try fancy stuff like debuginfo, which is useless on binary-only
# packages. Don't strip binary too
# Be sure buildpolicy set to do nothing
%define        __spec_install_post %{nil}
%define          debug_package %{nil}
%define        __os_install_post %{_dbpath}/brp-compress

Summary: A very simple toy bin rpm package
Name: pfcaramirezagent
Version: 1.0.5
Release: 1
License: GPL+
Group: Development/Tools
SOURCE0 : %{name}-%{version}.tar.gz
URL: http://admin.pfc.aramirez.es/

BuildRoot: /tmp/%{name}-%{version}-%{release}-root

%description
%{summary}

%prep
%setup -q

%build
# Empty section.

%install
rm -rf %{buildroot}
mkdir -p  %{buildroot}

# in builddir
cp -a * %{buildroot}

install -m 755 %{buildroot}/usr/lib/systemd/system/pfcaramirezagent.service /usr/lib/systemd/system/pfcaramirezagent.service

%clean
rm -rf %{buildroot}

%preun
if [ "$1" = "0" ]; then
    /sbin/service pfcaramirezagent stop
    /sbin/chkconfig pfcaramirezagent off
fi

%post
/sbin/chkconfig pfcaramirezagent on
/sbin/service pfcaramirezagent start

%files
%defattr(-,root,root,-)
/usr/lib/systemd/system/pfcaramirezagent.service
#%config(noreplace) %{_sysconfdir}/%{name}/%{name}.conf
%{_bindir}/*

%changelog