package com.yauritux.model.entity;

import java.math.BigDecimal;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.EnumType;
import javax.persistence.Enumerated;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.Table;

import com.yauritux.model.constant.JourneyType;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
@Entity
@Table(name = "fares")
public class Fare {

	@Id
	@GeneratedValue(strategy = GenerationType.AUTO)
	private Long id;
	
	@Column(name = "journey", nullable = false)
	private String journey;
	
	@Enumerated(EnumType.STRING)
	@Column(name = "journey_type", nullable = false)
	private JourneyType journeyType;
	
	/*
	@ManyToMany(mappedBy = "fares")
	private List<Zone> zones = new ArrayList<>();
	*/	
		
	@Column(name = "fare", nullable = false)
	private BigDecimal fare;
	
	public Long getId() {
		return id;
	}
	
	protected void setId(Long id) {
		this.id = id;
	}
	
	public String getJourney() {
		return journey;
	}
	
	public void setJourney(String journey) {
		this.journey = journey;
	}
	
	public JourneyType getJourneyType() {
		return journeyType;
	}
	
	public void setJourneyType(JourneyType journeyType) {
		this.journeyType = journeyType;
	}
	
	public BigDecimal getFare() {
		return fare;
	}
	
	public void setFare(BigDecimal fare) {
		this.fare = fare;
	}
	
	/*
	public List<Zone> getZones() {
		return zones;
	}
	
	public void setZones(List<Zone> zones) {
		this.zones = zones;
	}
	*/

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + ((id == null) ? 0 : id.hashCode());
		return result;
	}

	@Override
	public boolean equals(Object obj) {
		if (this == obj)
			return true;
		if (obj == null)
			return false;
		if (getClass() != obj.getClass())
			return false;
		Fare other = (Fare) obj;
		if (id == null) {
			if (other.id != null)
				return false;
		} else if (!id.equals(other.id))
			return false;
		return true;
	}
}